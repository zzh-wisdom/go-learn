package main

import (
	"encoding/base64"
	"fmt"
)
// Base64 可以使用 64 个可打印字符来表示二进制数据，电子邮件就是使用这种编码。
func main() {

	// 需要处理的字符串
	message := "Away from keyboard. https://golang.org/"

	// 编码消息
	encodedMessage := base64.StdEncoding.EncodeToString([]byte (message))

	// 输出编码完成的消息
	fmt.Println(encodedMessage)

	// 解码消息
	data, err := base64.StdEncoding.DecodeString(encodedMessage)

	// 出错处理
	if err != nil {
		fmt.Println(err)
	} else {
		// 打印解码完成的数据
		fmt.Println(string(data))
	}
}