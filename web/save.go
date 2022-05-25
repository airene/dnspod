package web

import (
	"dnspod-go/config"
	"dnspod-go/dns"
	"encoding/json"
	"net/http"
)

// Save 保存
func Save(writer http.ResponseWriter, request *http.Request) {

	conf, _ := config.GetConfigCache()
	_ = json.NewDecoder(request.Body).Decode(&conf)
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
