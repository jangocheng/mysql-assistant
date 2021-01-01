# business data model manager system

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


# 定义make变量
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
BINARY_PATH=./bin
BINARY_NAME=$(BINARY_PATH)/business_data_model
BINARY_UNIX=$(BINARY_NAME)_unix
CMD_BINARY_NAME=business_data_model_cmd
CMD_BINARY_UNIX=$(CMD_BINARY_NAME)_unix

# make 不指定动作时，默认执行第一个动作
default:build

# 定义build，test，clean，run，deps动作 和build-linux， docker-build动作
run:
	$(GO) run main.go
test:
	$(GOTEST) -v
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
cmd-clean:
	$(GOCLEAN)
	rm -f $(CMD_BINARY_NAME)
	rm -f $(CMD_BINARY_UNIX)
mod:
	$(GO) mod tidy

build: mod clean test
	$(GOBUILD) -o $(BINARY_NAME) -v main.go
	shasum -a 256 $(BINARY_NAME)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
	shasum -a 256 $(BINARY_NAME)
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
	shasum -a 256 $(BINARY_NAME)
build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
	shasum -a 256 $(BINARY_NAME)
docker-build: # 对于go，这个东西鸡肋没用， 没必要要docker中执行build
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/data_model_go golang:latest go build -o "$(BINARY_UNIX)" -v

# cmd 定义
cmd-build: mod cmd-clean
	$(GOBUILD) -o ./bin/$(CMD_BINARY_NAME) -v ./cmd/cmd.go

cmd-build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(CMD_BINARY_NAME) -v ./cmd/cmd.go
	shasum -a 256 $(CMD_BINARY_NAME)
cmd-build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(CMD_BINARY_NAME) -v ./cmd/cmd.go
	shasum -a 256 $(CMD_BINARY_NAME)
cmd-build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(CMD_BINARY_NAME) -v ./cmd/cmd.go
	shasum -a 256 $(CMD_BINARY_NAME)
cmd-docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/data_model_go golang:latest go build -o "$(CMD_BINARY_NAME)" -v

