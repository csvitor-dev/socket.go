package tcp

import (
	"fmt"
	"net"

	"github.com/csvitor-dev/socket.go/pkg/socket"
	"github.com/csvitor-dev/socket.go/pkg/types"
)

type TCPServer struct {
	listener *net.TCPListener
	handler types.Process
	address string
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
		address: address,
	}, nil
}

func (tcp *TCPServer) InstallProcess(process types.Process) {
	tcp.handler = process
}

func (tcp *TCPServer) ListenAndServe() {
	fmt.Printf("Server listen on '%v'\n", tcp.address)
	defer tcp.listener.Close()

	for {
		conn, err := tcp.listener.Accept()

		if err != nil {
			continue
		}
		
		go socket.ConnectionHandler(conn, tcp.handler)
	}
}
