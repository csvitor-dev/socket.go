package main

import (
	"fmt"

	"github.com/csvitor-dev/socket.go/pkg/utils"
	"github.com/csvitor-dev/socket.go/src/types/tcp"
)

func main() {
	client, err := tcp.NewTCPConnection("localhost:8080")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	for {
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
