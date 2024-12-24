package tcp

import (
	"log"
	"net"
)

type TCPConnection struct {
	connection *net.TCPConn
}

func (tcp *TCPConnection) NewTCPConnection(address string) error {
	tcpAddress, err := createTCPAddress(address)

	if err != nil {
		return err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddress)

	if err != nil {
		return err
	}
	tcp.connection = conn
	return nil
}

func (tcp *TCPConnection) Prepare(address string) {
	err := tcp.NewTCPConnection(address)

	if err != nil {
		log.Fatalln(err)
		return
	}
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