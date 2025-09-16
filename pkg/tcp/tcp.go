package tcp

import (
	"net"
)

func Server(port string) (net.Listener, error) {
	listener, err := net.Listen("tcp", port)
	return listener, err
}
