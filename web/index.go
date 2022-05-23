package web

import (
	"dnspod-go/config"
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//go:embed index.html
var indexEmbedFile embed.FS

const VersionEnv = "DNSPOD_GO_VERSION"

type writtingData struct {
	config.Config
	Version string
}

// Index 填写信息
func Index(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFS(indexEmbedFile, "index.html")
	if err != nil {
		fmt.Println("html 解析失败..")
		fmt.Println(err)
		return
	}

	conf, err := config.GetConfigCache()
	if err == nil {
		// 已存在配置文件
		_ = tmpl.Execute(writer, &writtingData{Config: conf, Version: os.Getenv(VersionEnv)})
		return
	}

	// 默认值
	if conf.Ipv4.URL == "" {
		conf.Ipv4.URL = "https://myip4.ipip.net, https://ip.3322.net"
	}

	_ = tmpl.Execute(writer, &writtingData{Config: conf, Version: os.Getenv(VersionEnv)})
}
