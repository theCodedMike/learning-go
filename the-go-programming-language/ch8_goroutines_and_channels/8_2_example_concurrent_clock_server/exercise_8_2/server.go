// 练习8.2：
// 实现一个并发FTP服务器。服务器应该解析客户端发来的一些命令，比如cd命令来切换目录，ls来列出目录内文件，get和send来传输文件，close来关闭连接。
// 你可以用标准的ftp命令来作为客户端，或者也可以自己实现一个。
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("accept error: ", err)
			continue
		}

		go handleConn(conn, listener)
	}
}

func handleConn(c net.Conn, l net.Listener) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		handleInput(c, strings.TrimSpace(input.Text()))
	}
	c.Close()
}

func handleInput(c net.Conn, cmd string) {
	writeContent := ""
	switch cmd {
	case "cd":
		writeContent = "cd successfully"
	case "ls":
		writeContent = "ls successfully"
	case "get":
		writeContent = "get successfully"
	case "send":
		writeContent = "send successfully"
	case "close":
		writeContent = "close successfully"
	default:
		writeContent = "unexpected cmd: " + cmd
	}
	if _, err := fmt.Fprintln(c, "\t", writeContent); err != nil {
		log.Print("print error: ", err)
	}
}
