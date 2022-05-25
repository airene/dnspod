package config

import (
	"dnspod-go/util"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

// Ipv4Reg IPv4正则
const Ipv4Reg = `((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])`

// Config 配置
type Config struct {
	Ipv4 Ipv4Config
	DNS  DNSConfig
	TTL  string `default:"600"`
}

// Ipv4Config IPv4配置
type Ipv4Config struct {
	URL     string
	Domains string
}

// DNSConfig DNS配置
type DNSConfig struct {
	ID     string
	Secret string
}

// ConfigCache ConfigCache
type cacheType struct {
	ConfigSingle *Config
	Err          error
	Lock         sync.Mutex
}

var cache = &cacheType{}

// GetConfigCache 获得配置
func GetConfigCache() (conf Config, err error) {
	cache.Lock.Lock()
	defer cache.Lock.Unlock()

	if cache.ConfigSingle != nil {
		return *cache.ConfigSingle, cache.Err
	}

	// init config
	cache.ConfigSingle = &Config{TTL: "600", Ipv4: Ipv4Config{URL: "https://myip4.ipip.net, https://ip.3322.net"}}

	configFilePath := util.GetConfigFilePath()
	_, err = os.Stat(configFilePath)
	if err != nil {
		cache.Err = err
		return *cache.ConfigSingle, err
	}

	byt, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Println("yaml配置文件读取失败")
		cache.Err = err
		return *cache.ConfigSingle, err
	}

	err = yaml.Unmarshal(byt, cache.ConfigSingle)
	if err != nil {
		log.Println("反序列化配置文件失败", err)
		cache.Err = err
		return *cache.ConfigSingle, err
	}
	// remove err
	cache.Err = nil
	return *cache.ConfigSingle, err
}

// SaveConfig 保存配置
func (conf *Config) SaveConfig() (err error) {
	cache.Lock.Lock()
	defer cache.Lock.Unlock()

	byt, err := yaml.Marshal(conf)
	if err != nil {
		log.Println(err)
		return err
	}

	configFilePath := util.GetConfigFilePath()
	err = ioutil.WriteFile(configFilePath, byt, 0600)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("配置文件已保存在: %s\n", configFilePath)

	// 清空配置缓存
	cache.ConfigSingle = nil

	return
}

// GetIpv4Addr 获得IPv4地址
func (conf *Config) GetIpv4Addr() (result string) {

	client := util.CreateHTTPClient()
	urls := strings.Split(conf.Ipv4.URL, ",")
	for _, url := range urls {
		url = strings.TrimSpace(url)
		resp, err := client.Get(url)
		if err != nil {
			log.Println(fmt.Sprintf("连接失败! <a target='blank' href='%s'>查看能否返回IP</a>,", url))
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("读取IPv4结果失败! 接口: ", url)
			continue
		}
		comp := regexp.MustCompile(Ipv4Reg)
		result = comp.FindString(string(body))
		if result != "" {
			return
		} else {
			log.Printf("获取IPv4结果失败! 接口: %s ,返回值: %s\n", url, result)
		}
	}

	return
}
