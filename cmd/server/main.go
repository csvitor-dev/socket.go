package main

import (
	"fmt"

	"github.com/csvitor-dev/go-socket/src/tcp"
)

func main() {
	server, err := tcp.NewTcpServer(":8080")

	if err != nil {
		fmt.Println(err)
		return
	}
	server.ListenAndServe()
}