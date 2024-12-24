package runner

import (
	"fmt"

	t "github.com/csvitor-dev/socket.go/src/types"
	"github.com/csvitor-dev/socket.go/utils"
)

func RunClient(client t.ClientConnection, address string) {
	for {
		t.MakeClientConnection(client, address)
		defer client.Close()

		text := utils.Input("Enter your text:")
		err := client.SendMessage([]byte(text))

		if err != nil {
			fmt.Println(err)
			return
		}
		message, err := client.RetriveMessage()

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Retrive: %v\n", string(message))
	}
}
