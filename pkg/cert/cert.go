package cert

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net"
	"strings"
	"time"
)

func ParseCert(crt string) (x509.Certificate, error) {
	certBlock, _ := pem.Decode([]byte(crt))
	if certBlock == nil {
		return x509.Certificate{}, errors.New("invalid PEM block")
	}
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return x509.Certificate{}, err
	}

	return *cert, nil
}

func ParseKey(key string) (crypto.Signer, error) {
	keyBlockDER, _ := pem.Decode([]byte(key))
	if keyBlockDER == nil {
		return nil, errors.New("invalid PEM block")
	}

	if keyBlockDER.Type != "PRIVATE KEY" && !strings.HasSuffix(keyBlockDER.Type, " PRIVATE KEY") {
		return nil, fmt.Errorf("unknown PEM header %q", keyBlockDER.Type)
	}

	if parse, err := x509.ParsePKCS1PrivateKey(keyBlockDER.Bytes); err == nil {
		return parse, nil
	}

	if parse, err := x509.ParsePKCS8PrivateKey(keyBlockDER.Bytes); err == nil {
		switch parse.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey, ed25519.PrivateKey:
			return parse.(crypto.Signer), nil
		default:
			return nil, fmt.Errorf("found unknown private key type in PKCS#8 wrapping: %T", key)
		}
	}

	if parse, err := x509.ParseECPrivateKey(keyBlockDER.Bytes); err == nil {
		return parse, nil
	}

	return nil, errors.New("解析私钥失败")
}

func EncodeCert(cert x509.Certificate) ([]byte, error) {
	pemCert := pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}
	return pem.EncodeToMemory(&pemCert), nil
}

func EncodeKey(key crypto.Signer) ([]byte, error) {
	var pemType string
	var keyBytes []byte
	switch key := key.(type) {
	case *ecdsa.PrivateKey:
		var err error
		pemType = "EC"
		keyBytes, err = x509.MarshalECPrivateKey(key)
		if err != nil {
			return nil, err
		}
	case *rsa.PrivateKey:
		pemType = "RSA"
		keyBytes = x509.MarshalPKCS1PrivateKey(key)
	case ed25519.PrivateKey:
		var err error
		pemType = "ED25519"
		keyBytes, err = x509.MarshalPKCS8PrivateKey(key)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("未知的密钥类型 %T", key)
	}
	pemKey := pem.Block{Type: pemType + " PRIVATE KEY", Bytes: keyBytes}
	return pem.EncodeToMemory(&pemKey), nil
}

// GenerateSelfSigned 生成自签名证书
func GenerateSelfSigned(names []string) (cert []byte, key []byte, err error) {
	// 生成根密钥对
	rootPublicKey, rootPrivateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	var ips []net.IP
	ip := false
	for _, item := range names {
		ipItem := net.ParseIP(item)
		if ipItem != nil {
			ip = true
			ips = append(ips, ipItem)
		}
	}

	rootTemplate := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "Rat Panel Root CA"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(25, 0, 0),
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign,
		SignatureAlgorithm:    x509.PureEd25519,
	}

	rootCertBytes, err := x509.CreateCertificate(rand.Reader, &rootTemplate, &rootTemplate, rootPublicKey, rootPrivateKey)
	if err != nil {
		return nil, nil, err
	}
	rootCertBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: rootCertBytes,
	}

	// 生成中间证书密钥对
	interPublicKey, interPrivateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	interTemplate := x509.Certificate{
		SerialNumber:          big.NewInt(2),
		Subject:               pkix.Name{CommonName: "Rat Panel CA"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 1, 0),
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		SignatureAlgorithm:    x509.PureEd25519,
	}

	interCertBytes, err := x509.CreateCertificate(rand.Reader, &interTemplate, &rootTemplate, interPublicKey, rootPrivateKey)
	if err != nil {
		return nil, nil, err
	}
	interCertBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: interCertBytes,
	}

	// 生成客户端证书密钥对
	clientPublicKey, clientPrivateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	clientTemplate := x509.Certificate{
		SerialNumber:       big.NewInt(3),
		Subject:            pkix.Name{CommonName: "Rat Panel"},
		NotBefore:          time.Now(),
		NotAfter:           time.Now().AddDate(1, 1, 0),
		KeyUsage:           x509.KeyUsageDigitalSignature,
		ExtKeyUsage:        []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		SignatureAlgorithm: x509.PureEd25519,
	}
	if ip {
		clientTemplate.IPAddresses = ips
	} else {
		clientTemplate.DNSNames = names
	}

	clientCertBytes, err := x509.CreateCertificate(rand.Reader, &clientTemplate, &interTemplate, clientPublicKey, interPrivateKey)
	if err != nil {
		return nil, nil, err
	}
	clientCertBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: clientCertBytes,
	}

	// 拼接证书链
	cert = append(cert, pem.EncodeToMemory(clientCertBlock)...)
	cert = append(cert, pem.EncodeToMemory(interCertBlock)...)
	cert = append(cert, pem.EncodeToMemory(rootCertBlock)...)

	// 编码私钥
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(clientPrivateKey)
	if err != nil {
		return nil, nil, err
	}
	key = pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	return cert, key, nil
}
