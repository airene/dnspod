.PHONY: build clean

VERSION=$(shell git describe --tags `git rev-list --tags --max-count=1`)
BIN=dnspod-go
DIR_SRC=.
DOCKER_CMD=docker

GO_ENV=CGO_ENABLED=0
GO_FLAGS=-ldflags="-X main.version=$(VERSION) -X 'main.buildTime=`date`' -extldflags -static" -trimpath
GO=$(GO_ENV) $(shell which go)
GOROOT=$(shell `which go` env GOROOT)
GOPATH=$(shell `which go` env GOPATH)

build: $(DIR_SRC)/main.go
	@$(GO) build $(GO_FLAGS) -o $(BIN) $(DIR_SRC)

build_docker_image:
	@$(DOCKER_CMD) build -f ./Dockerfile -t dnspod-go:$(VERSION) .

# clean all build result
clean:
	@$(GO) clean ./...
	@rm -f $(BIN)
	@rm -rf ./dist/*
