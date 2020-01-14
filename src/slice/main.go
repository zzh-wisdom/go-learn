// https://www.cnblogs.com/f-ck-need-u/p/9854932.html

package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	my_slice := make([]int,3,5)
	fmt.Println("my_slice len = ", len(my_slice))  // 3
	fmt.Println("my_slice cap = ", cap(my_slice))  // 5
	println("my_slice ", my_slice)      // [3/5]0xc42003df10
	// make只能构建slice、map和channel这3种结构的数据对象，
	// 因为它们都指向底层数据结构，都需要先为底层数据结构分配好内存并初始化。

	// 创建一个length和capacity都等于5的slice
	slice := make([]int,5)
	println("slice ", slice)

	// 创建长度为3的int数组
	array := [3]int{10, 20, 30}
	// 创建长度和容量都为3的slice
	slice = []int{10, 20, 30}
	println(cap(array))
	println("slice ", slice)

	// 声明一个nil slice
	// 这个slice不会指向哪个底层数组。也因此，nil slice的长度和容量都为0。
	var nil_slice []int
	println("nil_slice ", nil_slice)

	// 使用make创建
	// 空slice表示长度为0，容量为0，但却有指向的slice，只不过指向的底层数组暂时是长度为0的空数组。
	empty_slice := make([]int,0)
	// 直接创建
	//empty_slice := []int{}
	println("empty_slice ", empty_slice)

	// 当然，无论是nil slice还是empty slice，都可以对它们进行操作，如append()函数、len()函数和cap()函数。
	//+---------------------分割线----------------------+
	//对slice进行切片
	// 对slice切片的语法为：
	// SLICE[A:B]
	// SLICE[A:B:C]
	// 其中A表示从SLICE的第几个元素开始切，B控制切片长度的结束，左闭右开，C控制切片容量的结束(C-A)，
	// 如果没有给定C，则表示切到底层数组或者（被）切片的最尾部。
	slice = []int{10, 20, 30, 40, 50, 60}
	println("slice ", slice)
	slice1 := slice[2:3]
	println("slice1 ", slice1) //容量为4
	slice2 := slice1[1:3]
	println("slice2 ", slice2) //容量为3

	slice1 = slice[1:3:4]
	println("slice1 ", slice1) //容量为3
	slice2 = slice1[0:2]
	println("slice2 ", slice2) //容量为3

	// 测试int类型占几字节
	// 跟机器有关吧
	i := int(2);
	println(unsafe.Sizeof(i)); // 8
	j := 3;
	println(unsafe.Sizeof(j));

	// 还有几种简化形式：
	// SLICE[A:]  // 从A切到最尾部
	// SLICE[:B]  // 从最开头切到B(不包含B)
	// SLICE[:]   // 从头切到尾，等价于复制整个SLICE
	slice = []int{10, 20, 30}
	for i := range slice {
		println(i);
	}


}