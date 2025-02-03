package runner

import (
	"fmt"

	"github.com/csvitor-dev/socket.go/pkg/utils"
	t "github.com/csvitor-dev/socket.go/src/types"
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
