package udp

import "net"

type ClientConnection struct {
	connection *net.UDPConn
}

func NewClientConnection(address string) (ClientConnection, error) {
	udpAddress, err := CreateAddress(address)

	if err != nil {
		return ClientConnection{}, err
	}
	conn, err := net.DialUDP("udp4", nil, udpAddress)

	if err != nil {
		return ClientConnection{}, err
	}
	return ClientConnection{conn}, nil
}

func (c ClientConnection) SendMessage(message string) error {
	_, err := c.connection.Write([]byte(message))
	return err
}

func (c ClientConnection) RetriveMessage() (string, error) {
	buf := make([]byte, 1024)
	_, _, err := c.connection.ReadFromUDP(buf)

	return string(buf), err
}

func (c ClientConnection) Close() error {
	return c.connection.Close()
}