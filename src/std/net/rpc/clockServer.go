// clock1 是一个定期报告时间的 TCP 服务器
package main
import (
	"io"
	"log"
	"net"
	"time"
)
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) //例如，连接中止
			continue
		}
		handleConn(conn)   // 一次处理一个连接
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return //例如，连接断开
		}
		time.Sleep(1 * time.Second)
	}
}