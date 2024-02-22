package main

import (
	"io"
	"log"
	"net"
	"os"
)

// 在终端执行：
//
//	1.启动服务端
//	go run ./ch8_goroutines_and_channels/8_2_example_concurrent_clock_server/exercise_8_2/server.go
//
//	2.发起客户端请求连接
//	go run ./ch8_goroutines_and_channels/8_2_example_concurrent_clock_server/exercise_8_2/client.go
//
//	3. 输入命令，如cd、close等
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("dial error", err)
	}

	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
