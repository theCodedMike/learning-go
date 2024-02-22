// 练习8.1：
// 修改clock2来支持传入参数作为端口号，然后写一个clockwall的程序，这个程序可以同时与多个clock服务器通信，
// 从多个服务器中读取时间，并且在一个表格中一次显示所有服务器传回的结果，类似于你在某些办公室里看到的时钟墙。
// 如果你有地理学上分布式的服务器可以用的话，让这些服务器跑在不同的机器上面；或者在同一台机器上跑多个不同的实例，
// 这些实例监听不同的端口，假装自己在不同的时区。像下面这样：
// $ TZ=US/Eastern    ./clock2 -port 8010 &
// $ TZ=Asia/Tokyo    ./clock2 -port 8020 &
// $ TZ=Europe/London ./clock2 -port 8030 &
// $ clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	for _, arg := range os.Args[1:] {
		args := strings.SplitN(arg, "=", 2)

		if len(args) != 2 {
			fmt.Printf("Format Error: %s(e.g. NewYork=localhost:8010)\n", arg)
			continue
		}

		tz, addr := args[0], args[1]
		if !strings.Contains(addr, ":") {
			fmt.Printf("Format Error: %s(e.g. localhost:8010)\n", addr)
			continue
		}

		// 分别启动一个协程去处理监听任务
		go handleListen(tz, addr)
	}
	// 这里需要保证“主协程”一直存在，如果没有，那么main函数执行完后“主协程”就退出了，它不会等待“子协程”
	select {}
}

func handleListen(tz string, addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(tz, ", listen error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(tz, ", accept error:", err) // e.g., connection aborted
			continue
		}
		go handleConn(tz, conn) // handle connections concurrently
	}
}

func handleConn(tz string, c net.Conn) {
	defer c.Close()
	content := fmt.Sprintf("| %-15s | %-20s |\n", tz, time.Now().Format("15:04:05"))
	_, err := io.WriteString(c, content)
	if err != nil {
		log.Print(tz, ", write string error:", err)
	}
}
