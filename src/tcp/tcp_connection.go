package tcp

import (
	"fmt"
	"net"
)

type TCPConnection struct {
	connection *net.TCPConn
	address string
}

func NewTCPConnection(address string) (*TCPConnection, error) {
	tcpAddress, err := createTCPAddress(address)

	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddress)

	if err != nil {
		return nil, err
	}
	fmt.Printf("New client instance for '%v'\n", address)

	return &TCPConnection{connection: conn, address: address}, nil
}

func (tcp *TCPConnection) SendMessage(message []byte) error {
	_, err := tcp.connection.Write(message)
	return err
}

func (tcp *TCPConnection) RetriveMessage() ([]byte, error) {
	buf := make([]byte, 1024)
	_, err := tcp.connection.Read(buf)

	return buf, err
}

func (tcp *TCPConnection) Close() error {
	return tcp.connection.Close()
}