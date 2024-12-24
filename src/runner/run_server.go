package runner

import (
	"log"

	t "github.com/csvitor-dev/socket.go/src/types"
)

func RunServer(server t.ServerListener, address string) {
	t.MakeServerListener(server, address)
	
	log.Printf("Listen in <%v>\n", address)
	server.ListenAndServe()
}
