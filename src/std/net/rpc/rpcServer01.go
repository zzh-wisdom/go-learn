package main
import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)
// 算数运算结构体
type Arith struct {
}
// 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}
// 算数运算响应结构体
type ArithResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}
// 乘法运算方法
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}
// 除法运算方法
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除以零")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}
func main() {
	rpc.Register(new(Arith)) // 注册rpc服务
	rpc.HandleHTTP()         // 采用http协议作为rpc载体
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln("致命错误: ", err)
	}
	fmt.Fprintf(os.Stdout, "%s", "开始连接")
	http.Serve(lis, nil)
}