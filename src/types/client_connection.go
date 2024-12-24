package types

type ClientConnection interface {
	Prepare(string)
	SendMessage([]byte) error
	RetriveMessage() ([]byte, error)
	Close() error
}

func MakeClientConnection(client ClientConnection, address string) {
	client.Prepare(address)
}