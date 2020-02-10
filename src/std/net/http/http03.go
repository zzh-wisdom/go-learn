package main
import (
	"fmt"
	"net/http"
)
func main() {
	resp, err := http.Head("http://www.baidu.com")
	if err != nil {
		fmt.Println("Request Failed: ", err.Error())
		return
	}
	defer resp.Body.Close()
	// 打印头信息
	for key, value := range resp.Header {
		fmt.Println(key, ":", value)
	}
}