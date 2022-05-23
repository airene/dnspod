package config

import (
	"log"
	"strings"
)

// 固定的主域名
var staticMainDomains = []string{"com.cn", "org.cn", "net.cn", "ac.cn", "eu.org"}

// Domains Ipv4 domains
type Domains struct {
	Ipv4Addr    string
	Ipv4Domains []*Domain
}

// Domain 域名实体
type Domain struct {
	DomainName string
	SubDomain  string
}

// GetSubDomain 获得子域名，为空返回@
// 阿里云，dnspod需要
func (d Domain) GetSubDomain() string {
	if d.SubDomain != "" {
		return d.SubDomain
	}
	return "@"
}

// InitGetNewIp 接口获得ip并校验用户输入的域名 在程序启动时候执行
func (domains *Domains) InitGetNewIp(conf *Config) {
	domains.Ipv4Domains = checkParseDomains(conf.Ipv4.Domains)
	// IPv4
	if len(domains.Ipv4Domains) > 0 {
		ipv4Addr := conf.GetIpv4Addr()
		if ipv4Addr != "" {
			domains.Ipv4Addr = ipv4Addr
		} else {
			log.Println("未能获取IPv4地址, 将不会更新")
		}
	}

}

// checkParseDomains 校验并解析用户输入的域名
func checkParseDomains(domainArr []string) (domains []*Domain) {
	for _, domainStr := range domainArr {
		domainStr = strings.TrimSpace(domainStr)
		if domainStr != "" {
			domain := &Domain{}
			sp := strings.Split(domainStr, ".")
			length := len(sp)
			if length <= 1 {
				log.Println(domainStr, "域名不正确")
				continue
			}
			// 处理域名
			domain.DomainName = sp[length-2] + "." + sp[length-1]
			// 如包含在org.cn等顶级域名下，后三个才为用户主域名
			for _, staticMainDomain := range staticMainDomains {
				if staticMainDomain == domain.DomainName {
					domain.DomainName = sp[length-3] + "." + domain.DomainName
					break
				}
			}

			domainLen := len(domainStr) - len(domain.DomainName)
			if domainLen > 0 {
				domain.SubDomain = domainStr[:domainLen-1]
			} else {
				domain.SubDomain = domainStr[:domainLen]
			}

			domains = append(domains, domain)
		}
	}
	return
}
