package tcp

import (
	"net"
)

type ClientConnection struct {
	connection net.Conn
}

func NewClientConnection(address string) (ClientConnection, error) {
	conn, err := net.Dial("tcp", address)

	return ClientConnection{conn}, err
}

func (c ClientConnection) SendMessage(message string) error {
	_, err := c.connection.Write([]byte(message))
	return err
}

func (c ClientConnection) RetriveMessage() (string, error) {
	buf := make([]byte, 1024)
	_, err := c.connection.Read(buf)

	return string(buf), err
}

func (c ClientConnection) Close() error {
	return c.connection.Close()
}