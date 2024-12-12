package main

import (
	"github.com/csvitor-dev/go-socket/src/tcp"
)

func main() {
	tcp.RunServer(":8080")
}