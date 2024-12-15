package main

import (
	"github.com/csvitor-dev/socket.go/src/tcp"
)

func main() {
	tcp.RunServer("localhost:8080")
}