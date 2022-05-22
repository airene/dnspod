# ddns-go

<img src=https://img.shields.io/docker/image-size/airene/dnspod-go /> <img src=https://img.shields.io/docker/pulls/airene/dnspod-go />
持续获得你的公网IPv4，并解析到DNSPOD

## 动机

- `copy`自[ddns-go](https://github.com/jeessy2/ddns-go)项目
- 为了学习 golang、github actions、docker
- v1 在ddns-go的基础上做功能精简，只保留满足自己需要的， 主要是为了练手 2022-05-21
- v2开始做项目级别的rebuild和refactor

## 特性

- 只支持从能获得ip的网站获得ip地址后通过api修改在 **DNSPOD** 的域名解析

## Docker中使用

- 不挂载主机目录, 删除容器同时会删除配置

  ```bash
  # host模式
  docker run -d --name dnspod-go --restart=always --net=host airene/dnspod-go
  ```

- 在浏览器中打开`http://主机IP:9877`，修改你的配置，成功

挂载卷
  ```bash
  docker run -d --name dnspod-go --restart=always --net=host -v /opt/dnspod-go:/root airene/dnspod-go
  ```

- [可选] 支持启动带参数 `-l`监听地址 `-f`间隔时间(秒)

  ```bash
  docker run -d --name dnspod-go --restart=always --net=host airene/dnspod-go -l :9877 -f 600
  ```


## 界面

带补充

## 开发&自行编译

- 如果喜欢从源代码编译自己的版本，可以使用本项目提供的 Makefile 构建
- 开发环境 golang 1.17
- 使用 `make build` 生成本地编译后的 `dnspod-go` 可执行文件
- 使用 `make build_docker_image` 自行编译 Docker 镜像
