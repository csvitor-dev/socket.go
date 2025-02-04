package socket

import (
	"log"
	"net"

	"github.com/csvitor-dev/socket.go/pkg/types"
)

func ConnectionHandler(connection net.Conn, process types.Process) {
	defer connection.Close()

	request := make([]byte, 1024)
	_, err := connection.Read(request)
	
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("<RECEIVE>")
	
	response := process.Push(request)
	_, err = connection.Write(response)

	if err != nil {
		log.Fatalln(err)
		return
	}
	
	log.Println("<SEND>")
}