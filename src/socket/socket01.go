package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main(){
	addr, err:= net.ResolveTCPAddr("tcp", "127.0.0.1:8000")//端口号设置为0可以自动选择可用的端口号
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", addr) // 创建TCP4服务器端监听器
	if err != nil {
		log.Fatal(err) // Println + os.Exit(1)
	}
	fmt.Println(listener.Addr())//获取实际分配地址
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err) // 错误直接退出
		}
		fmt.Println("remote address:", conn.RemoteAddr())
		go echo(conn)
	}
}

func echo(conn *net.TCPConn) {
	tick := time.Tick(5 * time.Second) // 五秒的心跳间隔
	for now := range tick {
		n, err := conn.Write([]byte(now.String()))
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		fmt.Printf("send %d bytes to %s\n", n, conn.RemoteAddr())
	}
}
