package main

import (
	"fmt"
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
		row := 5
		col := 5
		for {
			<- ticker.C
			window.SetString("x", row, col, conn)
			row, col, count = Square(row, col, count)
			fmt.Printf("row: %v col: %v\n", row, col)
		}
	}
}

func Square(row, col, count int) (int, int, int) {
	count++

	if count <= 16 {
		row++
	} else if count > 16 && count <= 76 {
		col++
	} else if count > 76 && count <= 92 {
		row--
	} else if count > 92 && count <= 152 {
		col--
	} else {
		count = 0
	}

	fmt.Println("count: ", count)

	return row, col, count
}
