package main

import (
	"fmt"

	"github.com/csvitor-dev/socket.go/pkg/utils"
	"github.com/csvitor-dev/socket.go/src/tcp"
)

func main() {
	for {
		client, err := tcp.NewTCPConnection(":8080")

		if err != nil {
			fmt.Println(err)
			return
		}
		defer client.Close()

		text := utils.Input("Message: ")
		
		client.SendMessage([]byte(text))
	
		bytes, err := client.RetriveMessage()
	
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(bytes))
	}
}
