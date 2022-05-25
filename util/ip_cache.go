package util

// IpCache 上次IP缓存
type IpCache struct {
	Addr         string // 缓存地址
	ForceCompare bool   // 是否强制比对
}

var Ipv4Cache = &IpCache{}

func (d *IpCache) Check(newAddr string) bool {
	if newAddr == "" {
		return true
	}
	// 地址改变 或 强制比对
	if d.Addr != newAddr || d.ForceCompare {
		d.Addr = newAddr
		d.ForceCompare = false
		return true
	}
	d.Addr = newAddr
	return false
}
