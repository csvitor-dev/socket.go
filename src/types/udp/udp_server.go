package types

import (
	"log"
	"net"

	"github.com/csvitor-dev/socket.go/src/lib"
)

type UDPServer struct {
	connection *net.UDPConn
}

func createUDPAddress(address string) (*net.UDPAddr, error) {
	return net.ResolveUDPAddr("udp4", address)
}

func (udp *UDPServer) NewUDPServer(address string) error {
	addr, err := createUDPAddress(address)
	
	if err != nil {
		return err
	}
	conn, err := net.ListenUDP(address, addr)
	
	if err != nil {
		return err
	}
	udp.connection = conn
	return nil
}

func (udp *UDPServer) Prepare(address string) {
	err := udp.NewUDPServer(address)

	if err != nil {
		log.Fatalln(err)
		return
	}
}

func (udp *UDPServer) ListenAndServe() {
	for {
		lib.ConnectionHandler(udp.connection)
	}
}
	