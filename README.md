# redis-tool
适用于从远端redis根据匹配模式批量删除key和导出key的value

###linux

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o redis-tool main.go

###查看帮助

redis-tool -help
