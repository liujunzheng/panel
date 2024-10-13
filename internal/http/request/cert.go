package request

type CertCreate struct {
	Type      string   `form:"type" json:"type" validate:"required"`
	Domains   []string `form:"domains" json:"domains" validate:"required"`
	AutoRenew bool     `form:"auto_renew" json:"auto_renew"`
	AccountID uint     `form:"account_id" json:"account_id"`
	DNSID     uint     `form:"dns_id" json:"dns_id"`
	WebsiteID uint     `form:"website_id" json:"website_id"`
}

type CertUpdate struct {
	ID        uint     `form:"id" json:"id" validate:"required"`
	Type      string   `form:"type" json:"type" validate:"required"`
	Domains   []string `form:"domains" json:"domains" validate:"required"`
	AutoRenew bool     `form:"auto_renew" json:"auto_renew"`
	AccountID uint     `form:"account_id" json:"account_id"`
	DNSID     uint     `form:"dns_id" json:"dns_id"`
	WebsiteID uint     `form:"website_id" json:"website_id"`
}

type CertDeploy struct {
	ID        uint `form:"id" json:"id" validate:"required"`
	WebsiteID uint `form:"website_id" json:"website_id" validate:"required"`
}
