package config

import (
	"ddns-go/util"
	"log"
	"strings"
)

// 固定的主域名
var staticMainDomains = []string{"com.cn", "org.cn", "net.cn", "ac.cn", "eu.org"}

// 获取ip失败的次数
var getIPv4FailTimes = 0

// Domains Ipv4 domains
type Domains struct {
	Ipv4Addr    string
	Ipv4Domains []*Domain
	Ipv6Addr    string
	Ipv6Domains []*Domain
}

// Domain 域名实体
type Domain struct {
	DomainName string
	SubDomain  string
}

func (d Domain) String() string {
	if d.SubDomain != "" {
		return d.SubDomain + "." + d.DomainName
	}
	return d.DomainName
}

// GetFullDomain 获得全部的，子域名
func (d Domain) GetFullDomain() string {
	if d.SubDomain != "" {
		return d.SubDomain + "." + d.DomainName
	}
	return "@" + "." + d.DomainName
}

// GetSubDomain 获得子域名，为空返回@
// 阿里云，dnspod需要
func (d Domain) GetSubDomain() string {
	if d.SubDomain != "" {
		return d.SubDomain
	}
	return "@"
}

// GetNewIp 接口/网卡获得ip并校验用户输入的域名
func (domains *Domains) GetNewIp(conf *Config) {
	domains.Ipv4Domains = checkParseDomains(conf.Ipv4.Domains)
	// IPv4
	if len(domains.Ipv4Domains) > 0 {
		ipv4Addr := conf.GetIpv4Addr()
		if ipv4Addr != "" {
			domains.Ipv4Addr = ipv4Addr
			getIPv4FailTimes = 0
		} else {
			// 启用IPv4 & 未获取到IP & 填写了域名 & 失败刚好3次，防止偶尔的网络连接失败，并且只发一次
			getIPv4FailTimes++
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

// GetNewIpResult 获得GetNewIp结果
func (domains *Domains) GetNewIpResult(recordType string) (ipAddr string, retDomains []*Domain) {
	// IPv4
	if util.Ipv4Cache.Check(domains.Ipv4Addr) {
		return domains.Ipv4Addr, domains.Ipv4Domains
	} else {
		log.Printf("IPv4未改变，将等待 %d 次后与DNS服务商进行比对\n", util.MaxTimes-util.Ipv4Cache.Times+1)
		return "", domains.Ipv4Domains
	}
}
