package main

import "fmt"

func main()  {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)

	copy(a[2:], a[1:])
	fmt.Println(a)
}
