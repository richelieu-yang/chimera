APP=test

all: windows mac-m1 linux

#build:
#	@go build -tags=jsoniter -o ${APP} src/main.go
#	@echo "build命令结束"

mac-m1:
	@rm -rf $(APP)-darwin-arm64
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -tags=jsoniter -o $(APP)-darwin-arm64 test/test1.go

windows:
	@rm -rf $(APP)-windows-amd64.exe
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -tags=jsoniter -o $(APP)-windows-amd64.exe test/test1.go
	@#upx -9 $(APP)-windows-amd64.exe

linux:
	@rm -rf $(APP)-linux-amd64
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o $(APP)-linux-amd64 test/test1.go
	@#upx -9 $(APP)-linux-amd64
