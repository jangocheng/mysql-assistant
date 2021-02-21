# 项目数据模型管理系统

# 定义make变量
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
BINARY_PATH=./bin
CMD_BINARY_NAME=$(BINARY_PATH)/start_up
CMD_BINARY_UNIX=$(CMD_BINARY_NAME)_unix

# make 不指定动作时，默认执行第一个动作
default:build

# 定义build，test，clean，run，deps动作 和build-linux， docker-build动作
test:
	$(GOTEST) -v
clean:
	$(GOCLEAN)
	# rm -f $(BINARY_PATH)/*

mod:
	$(GO) mod tidy

build: mod clean test build-linux
	echo "build done"
#	$(GOBUILD) -o $(CMD_BINARY_NAME) -v ./cmd/cmd.go
#	shasum -a 256 $(CMD_BINARY_NAME)

# Cross compilation
build-linux:
	export CGO_ENABLED=0 GOOS=linux
	$(GOBUILD) -o $(CMD_BINARY_NAME)_linux -v ./cmd/cmd.go
	#shasum -a 256 $(CMD_BINARY_NAME)_linux

publish: clean-dir publish-linux

clean-dir:
	rm -rf ./release/* \!\(.gitkeep\)
	rm -rf ./release/.env
	rm -rf release_*.zip

publish-common-init:
	mkdir -p ./release/storage/logs && chmod -R 777 ./release/storage
	cp -r ./assets ./release
	cp .env.example ./release/.env
	cp ./business_event.sql ./release

publish-linux: publish-common-init
	cp $(CMD_BINARY_UNIX) ./release
	zip -r release_linux_`date +%Y%m%d`.zip release
