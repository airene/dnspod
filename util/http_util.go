package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

// GetHTTPResponse 处理HTTP结果，返回序列化的json
func GetHTTPResponse(resp *http.Response, url string, err error, result interface{}) error {
	body, err := getHTTPResponseOrg(resp, url, err)

	if err == nil {
		if len(body) != 0 {
			err = json.Unmarshal(body, &result)
			if err != nil {
				log.Printf("请求接口%s解析json结果失败! ERROR: %s\n", url, err)
			}
		}
	}

	return err

}

// GetHTTPResponseOrg 处理HTTP结果，返回byte
func getHTTPResponseOrg(resp *http.Response, url string, err error) ([]byte, error) {
	if err != nil {
		Ipv4Cache.ForceCompare = true
		log.Printf("请求接口%s失败! ERROR: %s\n", url, err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		Ipv4Cache.ForceCompare = true
		log.Printf("请求接口%s失败! ERROR: %s\n", url, err)
	}

	// 300及以上状态码都算异常
	if resp.StatusCode >= 300 {
		errMsg := fmt.Sprintf("请求接口 %s 失败! 返回内容: %s ,返回状态码: %d\n", url, string(body), resp.StatusCode)
		log.Println(errMsg)
		err = fmt.Errorf(errMsg)
	}

	return body, err
}

func CreateHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			IdleConnTimeout:     10 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}
