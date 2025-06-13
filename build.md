```shell
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o RemoteWebScreen.exe main.go
```
在macOS上跨平台需要：
brew install mingw-w64
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -ldflags="-s -w" -o RemoteWebScreen.exe main.go