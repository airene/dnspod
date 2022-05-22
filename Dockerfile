# build stage
FROM golang:1.17 AS builder

WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on \
    #&& go env -w GOPROXY=https://goproxy.cn,direct \
    && make clean build

# final stage
FROM alpine
LABEL name=dnspod-go
LABEL url=https://github.com/airene/dnspod-go

WORKDIR /app
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
#    && apk add --no-cache tzdata
RUN apk add --no-cache tzdata
#ENV TZ=Asia/Shanghai
COPY --from=builder /app/dnspod-go /app/dnspod-go
EXPOSE 9876
ENTRYPOINT ["/app/dnspod-go"]
CMD ["-l", ":9876", "-f", "300"]
