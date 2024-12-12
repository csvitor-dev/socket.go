package main

import (
	"github.com/csvitor-dev/go-socket/src/udp"
)

func main() {
	udp.RunClient("localhost:8080")
}
