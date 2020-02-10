package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(name)
	fmt.Println(os.Environ())
	fmt.Println("http_proxy:", os.Getenv("http_proxy"))
	fmt.Println(os.Getuid())
	fmt.Println(os.Getwd())
	f, err := exec.LookPath("main")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)

	u, _ := user.Current()
	log.Println("用户名：", u.Username)
	log.Println("用户id", u.Uid)
	log.Println("用户主目录：", u.HomeDir)
	log.Println("主组id：", u.Gid)
	// 用户所在的所有的组的id
	s, _ := u.GroupIds()
	log.Println("用户所在的所有组：", s)

	c := make(chan os.Signal, 0)
	signal.Notify(c)

	signal.Stop(c) //不允许继续往c中存入内容
	// Block until a signal is received.
	ss := <-c
	fmt.Println("Got signal:", ss)
}