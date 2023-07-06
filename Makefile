app=test2
path=test/test2.go

all: windows-amd64 darwin-arm64 linux-amd64 linux-arm64 linux-loong64 linux-mips64 linux-mips64le

darwin-arm64:
	@rm -rf $(app)-darwin-arm64
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -tags=jsoniter -o $(app)-darwin-arm64 $(path)

windows-amd64:
	@rm -rf $(app)-windows-amd64.exe
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -tags=jsoniter -o $(app)-windows-amd64.exe $(path)
	@upx -9 $(app)-windows-amd64.exe

linux-amd64:
	@rm -rf $(app)-linux-amd64
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o $(app)-linux-amd64 $(path)
	@upx -9 $(app)-linux-amd64

linux-arm64:
	@rm -rf $(app)-linux-arm64
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -tags=jsoniter -o $(app)-linux-arm64 $(path)
	@upx -9 $(app)-linux-arm64

linux-mips64:
	@rm -rf $(app)-linux-mips64
	@CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -tags=jsoniter -o $(app)-linux-mips64 $(path)

linux-mips64le:
	@rm -rf $(app)-linux-mips64le
	@CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -tags=jsoniter -o $(app)-linux-mips64le $(path)

linux-loong64:
	@rm -rf $(app)-linux-loong64
	@CGO_ENABLED=0 GOOS=linux GOARCH=loong64 go build -tags=jsoniter -o $(app)-linux-loong64 $(path)
