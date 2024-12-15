package types

type ClientConnection interface {
	SendMessage(string) error
	RetriveMessage() (string, error)
	Close() error
}
