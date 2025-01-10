# build stage
FROM golang:1.23 AS builder

WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on && make clean build

# final stage
FROM alpine
LABEL name=dnspod-go url=https://github.com/airene/dnspod-go
WORKDIR /app
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
COPY --from=builder /app/dnspod-go /app/dnspod-go
EXPOSE 9877
ENTRYPOINT ["/app/dnspod-go"]
CMD ["-l", ":9877", "-f", "3600"]
