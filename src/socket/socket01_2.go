package main
import (
	"log"
	"net"
)
func main() {
	//if len(os.Args) != 2 {
	//	log.Fatalf("Usage: %s host:port", os.Args[0])
	//}
	//service := os.Args[1]

	service := "127.0.0.1:8000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	n, err := conn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(n)
}