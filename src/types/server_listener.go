package types

type ServerListener interface {
	Prepare(string)
	ListenAndServe()
}

func MakeServerListener(server ServerListener, address string) {
	server.Prepare(address)
}
