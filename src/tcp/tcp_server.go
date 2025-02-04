package tcp

import (
	"net"

	"github.com/csvitor-dev/socket.go/pkg/socket"
	"github.com/csvitor-dev/socket.go/pkg/types"
)

type TCPServer struct {
	listener *net.TCPListener
	handler types.Process
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

	return &TCPServer{
		listener: listener,
		handler: &socket.DefaultHandler{},
	}, nil
}

func (tcp *TCPServer) InstallProcess(process types.Process) {
	tcp.handler = process
}

func (tcp *TCPServer) ListenAndServe() {
	for {
		conn, err := tcp.listener.Accept()

		if err != nil {
			continue
		}
		
		socket.ConnectionHandler(conn, tcp.handler)
	}
}
