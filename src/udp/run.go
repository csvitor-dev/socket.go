package udp

import "fmt"

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