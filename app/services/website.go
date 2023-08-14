// Package services 网站服务
package services

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/goravel/framework/facades"
	"golang.org/x/exp/slices"

	"panel/app/models"
	"panel/pkg/tools"
)

type Website interface {
	List(page int, limit int) (int64, []models.Website, error)
	Add(website PanelWebsite) (models.Website, error)
	Delete(id int) error
	GetConfig(id int) (WebsiteSetting, error)
}

type PanelWebsite struct {
	Name       string `json:"name"`
	Status     bool   `json:"status"`
	Domain     string `json:"domain"`
	Path       string `json:"path"`
	Php        int    `json:"php"`
	Ssl        bool   `json:"ssl"`
	Remark     string `json:"remark"`
	Db         bool   `json:"db"`
	DbType     string `json:"db_type"`
	DbName     string `json:"db_name"`
	DbUser     string `json:"db_user"`
	DbPassword string `json:"db_password"`
}

// WebsiteSetting 网站设置
type WebsiteSetting struct {
	Name              string   `json:"name"`
	Ports             []string `json:"ports"`
	Domains           []string `json:"domains"`
	Root              string   `json:"root"`
	Path              string   `json:"path"`
	Index             string   `json:"index"`
	Php               int      `json:"php"`
	OpenBasedir       bool     `json:"open_basedir"`
	Ssl               bool     `json:"ssl"`
	SslCertificate    string   `json:"ssl_certificate"`
	SslCertificateKey string   `json:"ssl_certificate_key"`
	HttpRedirect      bool     `json:"http_redirect"`
	Hsts              bool     `json:"hsts"`
	Waf               bool     `json:"waf"`
	WafMode           string   `json:"waf_mode"`
	WafCcDeny         string   `json:"waf_cc_deny"`
	WafCache          string   `json:"waf_cache"`
	Rewrite           string   `json:"rewrite"`
	Raw               string   `json:"raw"`
	Log               string   `json:"log"`
}

type WebsiteImpl struct {
	setting Setting
}

func NewWebsiteImpl() *WebsiteImpl {
	return &WebsiteImpl{
		setting: NewSettingImpl(),
	}
}

// List 列出网站
func (r *WebsiteImpl) List(page, limit int) (int64, []models.Website, error) {
	var websites []models.Website
	var total int64
	if err := facades.Orm().Query().Paginate(page, limit, &websites, &total); err != nil {
		return total, websites, err
	}

	return total, websites, nil
}

// Add 添加网站
func (r *WebsiteImpl) Add(website PanelWebsite) (models.Website, error) {
	// 禁止部分保留名称
	nameSlices := []string{"phpmyadmin", "mysql", "panel", "ssh"}
	if slices.Contains(nameSlices, website.Name) {
		return models.Website{}, errors.New("网站名称" + website.Name + "为保留名称，请更换")
	}

	// path为空时，设置默认值
	if len(website.Path) == 0 {
		website.Path = r.setting.Get(models.SettingKeyWebsitePath) + "/" + website.Name
	}
	// path不为/开头时，返回错误
	if website.Path[0] != '/' {
		return models.Website{}, errors.New("网站路径" + website.Path + "必须以/开头")
	}

	website.Ssl = false
	website.Status = true
	website.Domain = strings.TrimSpace(website.Domain)

	w := models.Website{
		Name:   website.Name,
		Status: website.Status,
		Path:   website.Path,
		Php:    website.Php,
		Ssl:    website.Ssl,
		Remark: website.Remark,
	}
	if err := facades.Orm().Query().Create(&w); err != nil {
		return w, err
	}

	tools.Mkdir(website.Path, 0755)

	index := `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="utf-8">
<title>耗子Linux面板</title>
</head>
<body>
<h1>耗子Linux面板</h1>
<p>这是耗子Linux面板的网站默认页面！</p>
<p>当您看到此页面，说明您的网站已创建成功。</p>
</body>
</html>
`
	tools.Write(website.Path+"/index.html", index, 0644)

	domainArr := strings.Split(website.Domain, "\n")
	portList := ""
	portArr := make(map[string]bool)
	domainList := ""
	for key, value := range domainArr {
		temp := strings.Split(value, ":")
		domainList += " " + temp[0]

		if len(temp) < 2 {
			if _, ok := portArr["80"]; !ok {
				if key == len(domainArr)-1 {
					portList += "    listen 80;"
				} else {
					portList += "    listen 80;\n"
				}
				portArr["80"] = true
			}
		} else {
			if _, ok := portArr[temp[1]]; !ok {
				if key == len(domainArr)-1 {
					portList += "    listen " + temp[1] + ";"
				} else {
					portList += "    listen " + temp[1] + ";\n"
				}
				portArr[temp[1]] = true
			}
		}
	}

	nginxConf := fmt.Sprintf(`# 配置文件中的标记位请勿随意修改，改错将导致面板无法识别！
# 有自定义配置需求的，请将自定义的配置写在各标记位下方。
server
{
    # port标记位开始
%s
    # port标记位结束
    # server_name标记位开始
    server_name%s;
    # server_name标记位结束
    # index标记位开始
    index index.php index.html;
    # index标记位结束
    # root标记位开始
    root %s;
    # root标记位结束

    # ssl标记位开始
    # ssl标记位结束

    # php标记位开始
    include enable-php-%d.conf;
    # php标记位结束

    # waf标记位开始
    waf on;
    waf_rule_path /www/server/openresty/ngx_waf/assets/rules/;
    waf_mode DYNAMIC;
    waf_cc_deny rate=1000r/m duration=60m;
    waf_cache capacity=50;
    # waf标记位结束

    # 错误页配置，可自行设置
    #error_page 404 /404.html;
    #error_page 502 /502.html;

    # 伪静态规则引入，修改后将导致面板设置的伪静态规则失效
    include /www/server/vhost/rewrite/%s.conf;

    # 面板默认禁止访问部分敏感目录，可自行修改
    location ~ ^/(\.user.ini|\.htaccess|\.git|\.svn)
    {
        return 404;
    }
    # 面板默认不记录静态资源的访问日志并开启1小时浏览器缓存，可自行修改
    location ~ .*\.(js|css)$
    {
        expires 1h;
        error_log /dev/null;
        access_log /dev/null;
    }

    access_log /www/wwwlogs/%s.log;
    error_log /www/wwwlogs/%s.log;
}
`, portList, domainList, website.Path, website.Php, website.Name, website.Name, website.Name)

	tools.Write("/www/server/vhost/"+website.Name+".conf", nginxConf, 0644)
	tools.Write("/www/server/vhost/rewrite/"+website.Name+".conf", "", 0644)
	tools.Write("/www/server/vhost/ssl/"+website.Name+".pem", "", 0644)
	tools.Write("/www/server/vhost/ssl/"+website.Name+".key", "", 0644)

	tools.Chmod(r.setting.Get(models.SettingKeyWebsitePath), 0755)
	tools.Chmod(website.Path, 0755)
	tools.Chown(r.setting.Get(models.SettingKeyWebsitePath), "www", "www")
	tools.Chown(website.Path, "www", "www")

	tools.Exec("systemctl reload openresty")

	rootPassword := r.setting.Get(models.SettingKeyMysqlRootPassword)
	if website.Db && website.DbType == "mysql" {
		tools.Exec(`/www/server/mysql/bin/mysql -uroot -p` + rootPassword + ` -e "CREATE DATABASE IF NOT EXISTS ` + website.DbName + ` DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;"`)
		tools.Exec(`/www/server/mysql/bin/mysql -uroot -p` + rootPassword + ` -e "CREATE USER '` + website.DbUser + `'@'localhost' IDENTIFIED BY '` + website.DbPassword + `';"`)
		tools.Exec(`/www/server/mysql/bin/mysql -uroot -p` + rootPassword + ` -e "GRANT ALL PRIVILEGES ON ` + website.DbName + `.* TO '` + website.DbUser + `'@'localhost';"`)
		tools.Exec(`/www/server/mysql/bin/mysql -uroot -p` + rootPassword + ` -e "FLUSH PRIVILEGES;"`)
	}

	return w, nil
}

// Delete 删除网站
func (r *WebsiteImpl) Delete(id int) error {
	var website models.Website
	if err := facades.Orm().Query().Where("id", id).FirstOrFail(&website); err != nil {
		return err
	}

	if _, err := facades.Orm().Query().Delete(&website); err != nil {
		return err
	}

	tools.Remove("/www/server/vhost/" + website.Name + ".conf")
	tools.Remove("/www/server/vhost/rewrite/" + website.Name + ".conf")
	tools.Remove("/www/server/vhost/ssl/" + website.Name + ".pem")
	tools.Remove("/www/server/vhost/ssl/" + website.Name + ".key")
	tools.Remove(website.Path)

	tools.Exec("systemctl reload openresty")

	return nil
}

// GetConfig 获取网站配置
func (r *WebsiteImpl) GetConfig(id int) (WebsiteSetting, error) {
	var website models.Website
	if err := facades.Orm().Query().Where("id", id).First(&website); err != nil {
		return WebsiteSetting{}, err
	}

	config := tools.Read("/www/server/vhost/" + website.Name + ".conf")

	var setting WebsiteSetting
	setting.Name = website.Name
	setting.Path = website.Path
	setting.Ssl = website.Ssl
	setting.Php = website.Php
	setting.Raw = config

	ports := tools.Cut(config, "# port标记位开始", "# port标记位结束")
	matches := regexp.MustCompile(`listen\s+(.*);`).FindAllStringSubmatch(ports, -1)
	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		setting.Ports = append(setting.Ports, match[1])
	}
	serverName := tools.Cut(config, "# server_name标记位开始", "# server_name标记位结束")
	match := regexp.MustCompile(`server_name\s+(.*);`).FindStringSubmatch(serverName)
	if len(match) > 1 {
		setting.Domains = strings.Split(match[1], " ")
	}
	root := tools.Cut(config, "# root标记位开始", "# root标记位结束")
	match = regexp.MustCompile(`root\s+(.*);`).FindStringSubmatch(root)
	if len(match) > 1 {
		setting.Root = match[1]
	}
	index := tools.Cut(config, "# index标记位开始", "# index标记位结束")
	match = regexp.MustCompile(`index\s+(.*);`).FindStringSubmatch(index)
	if len(match) > 1 {
		setting.Index = match[1]
	}

	if tools.Exists(setting.Root + "/.user.ini") {
		userIni := tools.Read(setting.Path + "/.user.ini")
		if strings.Contains(userIni, "open_basedir") {
			setting.OpenBasedir = true
		} else {
			setting.OpenBasedir = false
		}
	} else {
		setting.OpenBasedir = false
	}

	setting.SslCertificate = tools.Read("/www/server/vhost/ssl/" + website.Name + ".pem")
	setting.SslCertificateKey = tools.Read("/www/server/vhost/ssl/" + website.Name + ".key")
	if setting.Ssl {
		ssl := tools.Cut(config, "# ssl标记位开始", "# ssl标记位结束")
		setting.HttpRedirect = strings.Contains(ssl, "# http重定向标记位")
		setting.Hsts = strings.Contains(ssl, "# hsts标记位")
	} else {
		setting.HttpRedirect = false
		setting.Hsts = false
	}

	waf := tools.Cut(config, "# waf标记位开始", "# waf标记位结束")
	setting.Waf = strings.Contains(waf, "waf on;")
	match = regexp.MustCompile(`waf_mode\s+(.+);`).FindStringSubmatch(waf)
	if len(match) > 1 {
		setting.WafMode = match[1]
	}
	match = regexp.MustCompile(`waf_cc_deny\s+(.+);`).FindStringSubmatch(waf)
	if len(match) > 1 {
		setting.WafCcDeny = match[1]
	}
	match = regexp.MustCompile(`waf_cache\s+(.+);`).FindStringSubmatch(waf)
	if len(match) > 1 {
		setting.WafCache = match[1]
	}

	setting.Rewrite = tools.Read("/www/server/vhost/rewrite/" + website.Name + ".conf")
	setting.Log = tools.Escape(tools.Exec(`tail -n 100 '/www/wwwlogs/` + website.Name + `.log'`))

	return setting, nil
}
