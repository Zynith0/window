package tcphandler

import (
	"net"
)

func Echo(conn net.Conn, data string) {
	conn.Write([]byte(data + "\n"))
}

func EchoTrim(conn net.Conn, data string) {
	conn.Write([]byte(data))
}
