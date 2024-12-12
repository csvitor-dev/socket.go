package main

import (
	"github.com/csvitor-dev/go-socket/src/tcp"
)

func main() {
	tcp.RunClient("localhost:8080")
}
