package main

import (
	"github.com/csvitor-dev/go-socket/src/udp"
)

func main() {
	udp.RunServer("localhost:8080")
}