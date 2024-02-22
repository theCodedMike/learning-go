// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"io"
	"log"
	"net"
	"time"
)

// 以下服务端只能串行处理客户端的连接请求，无法做到同时处理多个请求
// 在终端执行：
//
//  1. 构建： go build ./ch8_goroutines_and_channels/8_2_example_concurrent_clock_server/clock1
//  2. 启动服务端： ./clock1 &
//  3. 在终端使用netcat命令发起TCP连接：nc localhost 8000
//
// 补充：
// 1. Ctrl + C 终止客户端连接
// 2. killall clock1 可以关闭服务端
func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
