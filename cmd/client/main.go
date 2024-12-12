package main

import (
	"fmt"

	"github.com/csvitor-dev/go-socket/src/tcp"
)

func main() {
	for {
		client, err := tcp.NewClientConnection("localhost:8080")

		if err != nil {
			fmt.Println(err)
			return
		}
		defer client.Close()

		text, err := input("Enter your text: ")
		if err != nil {
			fmt.Println(err)
			return
		}
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

func input(message string) (string, error) {
	fmt.Print(message)

    var text string
    _, err := fmt.Scanln(&text)

	return text, err
}