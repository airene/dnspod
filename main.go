package main

import (
	"dnspod-go/dns"
	"dnspod-go/web"
	"embed"
	"flag"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

// 监听地址 listen_port
var listen = flag.String("l", ":9877", "监听地址")

// 更新频率(秒) frequency
var every = flag.Int("f", 3600, "同步间隔时间(秒)")

//go:embed static
var staticEmbededFiles embed.FS

// version
var version = "DEV"

func main() {
	flag.Parse()
	if _, err := net.ResolveTCPAddr("tcp", *listen); err != nil {
		log.Fatalf("解析监听地址异常，%s", err)
	}
	_ = os.Setenv(web.VersionEnv, version)

	// 延时10秒运行
	run(10 * time.Second)

}

func run(firstDelay time.Duration) {
	// 启动静态文件服务
	http.Handle("/", http.FileServer(getFileSystem()))

	http.HandleFunc("/config", web.Config)
	http.HandleFunc("/save", web.Save)
	http.HandleFunc("/logs", web.Logs)
	http.HandleFunc("/clearLog", web.ClearLog)

	log.Println("监听端口", *listen, "...")

	// 定时运行
	go dns.RunTimer(firstDelay, time.Duration(*every)*time.Second)
	err := http.ListenAndServe(*listen, nil)

	if err != nil {
		log.Println("启动端口发生异常, 请检查端口是否被占用", err)
		time.Sleep(time.Minute)
		os.Exit(1)
	}
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(staticEmbededFiles, "static")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(fsys)
}
