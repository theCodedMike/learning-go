package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// 在终端执行：
//
//	1.启动服务端
//	go run ./ch8_goroutines_and_channels/8_2_example_concurrent_clock_server/exercise_8_1/server.go NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
//
//	2.发起客户端请求连接
//	go run ./ch8_goroutines_and_channels/8_2_example_concurrent_clock_server/exercise_8_1/client.go
func main() {
	addrs := []string{"localhost:8010", "localhost:8020", "localhost:8030"}
	fmt.Printf("| %-15s | %-20s |\n", "City", "Time")
	fmt.Printf("|-%15s-|-%20s-|\n", strings.Repeat("-", 15), strings.Repeat("-", 20))
	for _, addr := range addrs {
		handleDial(addr)
	}
}

func handleDial(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Fatal(err)
	}
}
