package udp

import (
	"fmt"

	"github.com/csvitor-dev/socket.go/utils"
)

func RunServer(address string) {
	server, err := NewUdpServer(address)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Listen on port: '%v'\n", address)
	server.ListenAndServe()
}

func RunClient(address string) {
	for {
		client, err := NewClientConnection(address)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer client.Close()
		text := utils.Input("Enter your text: ")
		err = client.SendMessage(text)

		if err != nil {
			fmt.Println(err)
			return
		}
		message, err := client.RetriveMessage()

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Retrive: %v\n", message)
	}
}
