# 项目数据模型管理系统
# 出于操作友好性，mac windows linux 三个makefile 独立分开， make指定文件构建指定平台的程序。
# 默认makefile 就针对下前机器构建就好。  其他开发者构建时，不会多余构建。

# 参考
# 在 Golang 中开发中使用 Makefile
# https://studygolang.com/articles/11131

# Golang的跨平台编译程序
# https://www.cnblogs.com/ghj1976/archive/2013/04/19/3030703.html
#各平台的GOOS和GOARCH参考
#OS                   ARCH                          OS version
#linux                386 / amd64 / arm             >= Linux 2.6
#darwin               386 / amd64                   OS X (Snow Leopard + Lion)
#freebsd              386 / amd64                   >= FreeBSD 7
#windows              386 / amd64                   >= Windows 2000

# Golang 编译Mac、Linux、Windows多平台可执行程序
# https://studygolang.com/articles/28339?fr=sidebar

# Docker命令_各种参数简介（run、v、rm、-w、-u、-e）
# https://blog.csdn.net/sxzlc/article/details/107676425

# Makefile 入门
# https://zhuanlan.zhihu.com/p/149346441

# 使用xgo编译支持CGO

# 定义make变量
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
BINARY_PATH=./bin
CMD_BINARY_NAME=$(BINARY_PATH)/start_up

# make 不指定动作时，默认执行第一个动作
default:build

test:
	$(GOTEST) -v
clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH)/*

mod:
	$(GO) mod tidy

build: mod clean test build-local
	echo "build done"
#	$(GOBUILD) -o $(CMD_BINARY_NAME) -v ./cmd/cmd.go
#	shasum -a 256 $(CMD_BINARY_NAME)

build-local:
	export CGO_ENABLED=0
	$(GOBUILD) -o $(CMD_BINARY_NAME) -v ./cmd/cmd.go
	#shasum -a 256 $(CMD_BINARY_NAME)_mac
build-by-docker:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/data_model_go golang:latest go build -o "$(CMD_BINARY_NAME)" -v

publish: clean-dir publish-local

clean-dir:
	rm -rf ./release/* \!\(.gitkeep\)
	rm -rf ./release/.env
	rm -rf release_*.zip

publish-common-init:
	mkdir -p ./release/storage/logs && chmod -R 777 ./release/storage
	cp -r ./assets ./release
	cp .env.example ./release/.env
	cp ./business_event.sql ./release

publish-local: publish-common-init
	cp $(CMD_BINARY_NAME) ./release
	zip -r release_`date +%Y%m%d`.zip release
	rm -rf ./release/* \!\(.gitkeep\)
	test -e ./release/.env && rm ./release/.env

create-docker-image:

