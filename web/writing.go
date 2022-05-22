package web

import (
	"ddns-go/config"
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//go:embed writing.html
var writingEmbedFile embed.FS

const VersionEnv = "DNSPOD_GO_VERSION"

type writtingData struct {
	config.Config
	Version string
}

// Writing 填写信息
func Writing(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFS(writingEmbedFile, "writing.html")
	if err != nil {
		fmt.Println("Error happened..")
		fmt.Println(err)
		return
	}

	conf, err := config.GetConfigCache()
	if err == nil {
		// 已存在配置文件，隐藏真实的ID、Secret
		tmpl.Execute(writer, &writtingData{Config: conf, Version: os.Getenv(VersionEnv)})
		return
	}

	// 默认值
	if conf.Ipv4.URL == "" {
		conf.Ipv4.URL = "https://myip4.ipip.net, https://ip.3322.net"
	}

	// 默认禁止外部访问
	conf.NotAllowWanAccess = true

	tmpl.Execute(writer, &writtingData{Config: conf, Version: os.Getenv(VersionEnv)})
}
