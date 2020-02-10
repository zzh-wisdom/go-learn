package main
import "fmt"
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func rawPrint(rawList ...interface{}) {

	// 遍历可变参数切片
	for _, a := range rawList {

		// 打印参数
		fmt.Println(a)
	}
}

// 打印函数封装
func print(slist ...interface{}) {

	// 将slist可变参数切片完整传递给下一个函数
	rawPrint(slist...)
	//rawPrint("fmt", slist)
}

func main() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)

	print(1, 2, 3)
}