package main

import (
	"fmt"
	// "strings"
	"time"
	"github.com/Zynith0/window/internal/tcpHandler"
	"github.com/Zynith0/window/pkg/tcp"
	"github.com/Zynith0/window/pkg/window"
)

func main() {
	server, err := tcp.Server(":8080")
	if err != nil {
		fmt.Println("skill issue", err)
	}
	defer server.Close()

	win := window.NewWindow(80, 24)

	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("skill issue", err)
			continue
		}
		tcphandler.Echo(conn, win.Cmd)

		count := 0
		for {
			<- ticker.C
			winString := window.SetString("x", 3, count, conn)
			tcphandler.Echo(conn, winString.String)
			fmt.Println(winString.String)
			count++
		}
	}
}
