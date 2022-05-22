package dns

import (
	"ddns-go/config"
	"log"
	"time"
)

// DNS interface
type DNS interface {
	Init(conf *config.Config)
	AddUpdateDomainRecords() (domains config.Domains)
}

// RunTimer 定时运行
func RunTimer(firstDelay time.Duration, delay time.Duration) {
	log.Printf("第一次运行将等待 %d 秒后运行 (等待网络)", int(firstDelay.Seconds()))
	time.Sleep(firstDelay)
	for {
		RunOnce()
		time.Sleep(delay)
	}
}

// RunOnce RunOnce
func RunOnce() {
	conf, err := config.GetConfigCache()
	if err != nil {
		return
	}

	var dnsImpl DNS = &Dnspod{}
	dnsImpl.Init(&conf)
	dnsImpl.AddUpdateDomainRecords()
}
