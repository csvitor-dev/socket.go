package main

import (
	"github.com/csvitor-dev/socket.go/src/tcp"
)

func main() {
	tcp.RunClient("localhost:8080")
}
