package main

import "fmt"

const (
	FlagNone = 1 << iota
	FlagRed
	FlagGreen
	FlagBlue
)

// 声明芯片类型
type ChipType int

const (
	None ChipType = iota
	CPU    // 中央处理器
	GPU    // 图形处理器
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}

	return "N/A"
}

func main()  {
	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue)

	// 输出CPU的值并以整型格式显示
	fmt.Printf("%s %d %T\n", CPU, CPU, CPU)
	fmt.Printf("%s %d", ChipType(2), 2)
}

