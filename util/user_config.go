package util

import (
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

const ConfigFilePathENV = "DNSPOD_CONFIG_FILE_PATH"

// GetConfigFilePath 获得配置文件路径
func GetConfigFilePath() string {
	configFilePath := os.Getenv(ConfigFilePathENV)
	if configFilePath != "" {
		return configFilePath
	}
	return getConfigFilePathDefault()
}

// 获得默认的配置文件路径
func GetConfigFilePathDefault() string {
	dir, err := homedir.Dir()
	if err != nil {
		log.Println("Geting current user failed!")
		return "../.dnspod_go_config.yaml"
	}
	return dir + string(os.PathSeparator) + ".dnspod_go_config.yaml"
}
