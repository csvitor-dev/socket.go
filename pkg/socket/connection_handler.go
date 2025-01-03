package socket

import (
	"fmt"
	"net"
	"strings"

	"github.com/csvitor-dev/socket.go/pkg/utils"
)

func ConnectionHandler(connection net.Conn) bool {
	defer connection.Close()

	buf := make([]byte, 1024)
	_, err := connection.Read(buf)

	if err != nil {
		return false
	}
	message := strings.ToUpper(string(buf))
	_, err = connection.Write([]byte(message))

	if err != nil {
		return false
	}
	fmt.Printf("\n[%v] >> %v \n", utils.LogDate(), message)

	return message == "EXIT"
}