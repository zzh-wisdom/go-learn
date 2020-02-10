package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	var str1 = "123";
	var str2 = "122";
	fmt.Println(str1==str2);

	const str = `第一行
第二行
第三行
\r\n`
	fmt.Println(str)
	tip := "1";
	tip1 := "忍着1"
	fmt.Printf("%s  len:%d, count:%d\n", tip, len(tip), utf8.RuneCountInString(tip))
	fmt.Printf("%s  len:%d, count:%d\n", tip1, len(tip1), utf8.RuneCountInString(tip1))

	hammer := "吃我一锤"

	sickle := "死吧"

	// 声明字节缓冲
	var stringBuilder bytes.Buffer

	// 把字符串写入缓冲
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)

	// 将缓冲以字符串形式输出
	fmt.Println(stringBuilder.String())
}

