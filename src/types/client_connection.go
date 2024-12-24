package types

type ClientConnection interface {
	SendMessage([]byte) error
	RetriveMessage() ([]byte, error)
	Close() error
}
