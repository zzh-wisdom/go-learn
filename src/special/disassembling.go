package main

func swap (a, b int) (x int, y int) {
	x = b
	y = a
	return
}

func main() {
	swap(10, 20)
}