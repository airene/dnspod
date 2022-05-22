package web

import (
	"ddns-go/config"
	"ddns-go/dns"
	"ddns-go/util"
	"net/http"
	"strings"
)

// Save 保存
func Save(writer http.ResponseWriter, request *http.Request) {

	conf, _ := config.GetConfigCache()

	idNew := request.FormValue("DnsID")
	secretNew := request.FormValue("DnsSecret")

	conf.DNS.ID = idNew
	conf.DNS.Secret = secretNew

	// 覆盖以前的配置

	conf.Ipv4.URL = strings.TrimSpace(request.FormValue("Ipv4Url"))
	conf.Ipv4.Domains = strings.Split(request.FormValue("Ipv4Domains"), "\r\n")

	conf.NotAllowWanAccess = request.FormValue("NotAllowWanAccess") == "on"
	conf.TTL = request.FormValue("TTL")

	// 保存到用户目录
	err := conf.SaveConfig()

	// 只运行一次
	util.Ipv4Cache.ForceCompare = true
	go dns.RunOnce()

	// 回写错误信息
	if err == nil {
		writer.Write([]byte("ok"))
	} else {
		writer.Write([]byte(err.Error()))
	}

}
