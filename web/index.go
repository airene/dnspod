package web

import (
	"dnspod-go/config"
	"encoding/json"
	"net/http"
	"os"
)

const VersionEnv = "DNSPOD_GO_VERSION"

type writtingData struct {
	config.Config
	Version string
}

// Config 填写信息
func Config(writer http.ResponseWriter, request *http.Request) {

	conf, _ := config.GetConfigCache()
	writer.Header().Set("Content-Type", "application-json")
	result, _ := json.Marshal(&writtingData{Config: conf, Version: os.Getenv(VersionEnv)})
	_, _ = writer.Write(result)

}
