package udp

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type UdpServer struct {
	connection *net.UDPConn
}

func CreateAddress(address string) (*net.UDPAddr, error) {
	udpAddress, err := net.ResolveUDPAddr("udp4", address)

	if err != nil {
		return nil, err
	}
	return udpAddress, nil
}

func NewUdpServer(address string) (UdpServer, error) {
	udpAddress, err := CreateAddress(address)

	if err != nil {
		return UdpServer{}, err
	}
	conn, err := net.ListenUDP("udp4", udpAddress)

	if err != nil {
		return UdpServer{}, err
	}
	return UdpServer{conn}, nil
}

func (s UdpServer) ListenAndServe() {
	defer s.connection.Close()

	for {
		buf := make([]byte, 1024)
		_, address, err := s.connection.ReadFromUDP(buf)

		if err != nil {
			continue
		}
		message := strings.ToUpper(string(buf))
		_, err = s.connection.WriteToUDP([]byte(message), address)

		if err != nil {
			return
		}

		fmt.Printf("\n[%v] >> %v \n", date(), message)
	}
}

func date() string {
	now := time.Now()
	year, month, day := now.Date()
	hour, min, sec := now.Clock()

	return fmt.Sprintf("%v/%v/%v | %v:%v:%v", year, month, day,
		 hour, min, sec)
}		