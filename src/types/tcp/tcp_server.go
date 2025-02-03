package tcp

import (
	"net"

	"github.com/csvitor-dev/socket.go/pkg/socket"
)

type TCPServer struct {
	listener *net.TCPListener
}

func createTCPAddress(address string) (*net.TCPAddr, error) {
	return net.ResolveTCPAddr("tcp", address)
}

func NewTCPServer(address string) (*TCPServer, error) {
	addr, err := createTCPAddress(address)
	
	if err != nil {
		return nil, err
	}
	listener, err := net.ListenTCP("tcp", addr)
	
	if err != nil {
		return nil, err
	}

	return &TCPServer{listener: listener}, nil
}

func (tcp *TCPServer) ListenAndServe() {
	for {
		conn, err := tcp.listener.Accept()

		if err != nil {
			continue
		}
		
		socket.ConnectionHandler(conn)
	}
}
