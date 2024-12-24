package main

import (
	"github.com/csvitor-dev/socket.go/src/runner"
	"github.com/csvitor-dev/socket.go/src/types/tcp"
)

func main() {
	runner.RunServer(&tcp.TCPServer{}, "localhost:8080")
}