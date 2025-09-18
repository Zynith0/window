package window

import (
	"net"
	"strconv"

	"github.com/Zynith0/window/internal/tcpHandler"
)

type Window struct {
	Cmd string
	String string
}

func NewWindow(intWidth, intHeight int) Window {
	var win Window

	width := strconv.Itoa(intWidth)
	height := strconv.Itoa(intHeight)

	win.Cmd = "<cmd>lua local buf=vim.api.nvim_create_buf(false, false); local max_height=vim.api.nvim_win_get_height(0); local max_width=vim.api.nvim_win_get_width(0); vim.api.nvim_open_win(buf, true, {relative='editor', width=" + width + ", height=" + height + ", col=(max_width - " + width + ") / 2, row=(max_height - " + height + ") / 2})<CR>"

	return win
}

func SetString(str string, rowInt, colInt int, conn net.Conn) {
	var win Window

	row := strconv.Itoa(rowInt)
	col := strconv.Itoa(colInt)

	win.String = "<cmd>lua vim.api.nvim_buf_set_text(0, " +  row + ", " +  col + ", " + row + ", " + col + ", " + "{ " + "'" + str + "'" + " }" + ")<CR>"

	tcphandler.Echo(conn, win.String)
}
