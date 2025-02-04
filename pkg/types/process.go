package types

type Process interface {
	Push([]byte) []byte
}