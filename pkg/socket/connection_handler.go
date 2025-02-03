package socket

import (
	"log"
	"net"

	"github.com/csvitor-dev/socket.go/pkg/utils"
)

func ConnectionHandler(connection net.Conn) {
	defer connection.Close()

	buf := make([]byte, 1024)
	_, err := connection.Read(buf)

	log.Printf("[%v] <RECEIVE>\n", utils.LogDate())
	
	if err != nil {
		return
	}
	_, _ = connection.Write(buf)
	
	log.Printf("[%v] <SEND>\n", utils.LogDate())
}