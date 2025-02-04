package main

import (
	"fmt"

	"github.com/csvitor-dev/socket.go/src/tcp"
)

func main() {
	server, err := tcp.NewTCPServer("localhost:8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	server.ListenAndServe()
}