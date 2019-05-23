myredis:
	GOOS=linux GOARCH=amd64 CGO_ENABLE=0 go build -o myredis main.go
