# 项目数据模型管理系统
# 现在需要解决的问题是怎么让xgo docker 走仓库镜像代理

# 定义make变量
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
BINARY_PATH=./bin
#CMD_BINARY_NAME=$(BINARY_PATH)/start_up
#CMD_BINARY_WINDOW=$(CMD_BINARY_NAME)

# make 不指定动作时，默认执行第一个动作
default:build

# 定义build，test，clean，run，deps动作 和build-linux， docker-build动作
test:
	$(GOTEST) -v
clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH)/*

mod:
	$(GO) mod tidy

build: mod clean test build-windows
	echo "build done"
#	$(GOBUILD) -o $(CMD_BINARY_NAME) -v ./cmd/cmd.go
#	shasum -a 256 $(CMD_BINARY_NAME)

# Cross compilation
build-windows:
	xgo --image="youwen21/ali-proxy-xgo" --targets="windows/*" -dest=$(BINARY_PATH) ./cmd/
	#export CGO_ENABLED=0 GOOS=windows GOARCH=386
	#$(GOBUILD) -o $(CMD_BINARY_NAME)_windows.exe -v ./cmd/cmd.go
	#shasum -a 256 $(CMD_BINARY_NAME)_windows

publish: clean-dir publish-windows

clean-dir:
	rm -rf ./release/* \!\(.gitkeep\)
	rm -rf ./release/.env
	rm -rf release_*.zip

publish-common-init:
	mkdir -p ./release/storage/logs && chmod -R 777 ./release/storage
	cp -r ./assets ./release
	cp .env.example ./release/.env
	cp ./business_event.sql ./release

publish-windows: publish-common-init
	cp $(BINARY_PATH)/*.exe ./release
	zip -r release_windows_`date +%Y%m%d`.zip release
