package web

import (
	"dnspod-go/config"
	"dnspod-go/dns"
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

	conf.Ipv4.URL = strings.TrimSpace(request.FormValue("Ipv4Url"))
	conf.Ipv4.Domains = strings.Split(request.FormValue("Ipv4Domains"), "\r\n")
	conf.TTL = request.FormValue("TTL")

	// 保存到用户目录
	err := conf.SaveConfig()

	go dns.RunOnce()

	// 回写错误信息
	if err == nil {
		_, _ = writer.Write([]byte("ok"))
	} else {
		_, _ = writer.Write([]byte(err.Error()))
	}

}
