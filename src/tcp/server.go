package tcp

import (
	"fmt"
	"net"
	"strings"
)

type TcpServer struct {
	listener net.Listener
}

func NewTcpServer(address string) (TcpServer, error) {
	listener, err := net.Listen("tcp", address)

	return TcpServer{listener}, err
}

func (s TcpServer) ListenAndServe() {
	for {
		conn, err := s.listener.Accept()

		if err == nil {
			handler(conn)
		}
	}
}

func handler(connection net.Conn) {
	defer connection.Close()

	buf := make([]byte, 1024)
	_, err := connection.Read(buf)

	if err != nil {
		return
	}
	message := strings.ToUpper(string(buf))
	_, err = connection.Write([]byte(message))

	if err != nil {
		return
	}
	fmt.Printf("Message: %v\n", message)
}