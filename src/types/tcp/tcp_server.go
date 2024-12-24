package tcp

import (
	"log"
	"net"

	"github.com/csvitor-dev/socket.go/src/lib"
)

type TCPServer struct {
	listener *net.TCPListener
}

func createTCPAddress(address string) (*net.TCPAddr, error) {
	return net.ResolveTCPAddr("tcp", address)
}

func (tcp *TCPServer) NewTCPServer(address string) error {
	addr, err := createTCPAddress(address)
	
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP(address, addr)
	
	if err != nil {
		return err
	}
	tcp.listener = listener
	return nil
}

func (tcp *TCPServer) Prepare(address string) {
	err := tcp.NewTCPServer(address)

	if err != nil {
		log.Fatalln(err)
		return
	}
}

func (tcp *TCPServer) ListenAndServe() {
	for {
		conn, err := tcp.listener.Accept()

		if err != nil {
			continue
		}
		
		if ok := lib.HandlerConnection(conn); ok {
			return
		}
	}
}
