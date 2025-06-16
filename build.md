```shell
# Windows 32位
GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o HTRS.exe main.go
# Windows 64位
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o HTRS.exe main.go
```
注意：要支持 Windows 7 的话 Go 版本最高只能是 1.20

目前的问题：
多客户端：多个客户端连接时，变化只会推送到一个
重连：无法连接时会报错，不会继续重连