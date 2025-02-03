package main

import (
	"fmt"

	"github.com/csvitor-dev/socket.go/src/types/tcp"
)

func main() {
	server, err := tcp.NewTCPServer(":8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	server.ListenAndServe()
}