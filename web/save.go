package web

import (
	"dnspod-go/config"
	"dnspod-go/dns"
	"dnspod-go/util"
	"encoding/json"
	"net/http"
)

// Save 保存
func Save(writer http.ResponseWriter, request *http.Request) {

	conf, _ := config.GetConfigCache()
	_ = json.NewDecoder(request.Body).Decode(&conf)
	// 保存到用户目录
	err := conf.SaveConfig()
	// 有配置修改强行检查dns记录
	util.Ipv4Cache.ForceCompare = true
	go dns.RunOnce()
	// 回写错误信息
	if err == nil {
		_, _ = writer.Write([]byte("ok"))
	} else {
		_, _ = writer.Write([]byte(err.Error()))
	}

}
