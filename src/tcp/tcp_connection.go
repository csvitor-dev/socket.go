package tcp

import (
	"net"
)

type TCPConnection struct {
	connection *net.TCPConn
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

	return &TCPConnection{connection: conn}, nil
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