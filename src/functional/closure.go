package main

import "fmt"

// GO 语言的闭包展示

func main() {

	result := closureDemo()(5)
	fmt.Printf("闭包计算结果：%v\n", result)
}

func closureDemo() func(int) int {
	baseNumber := 4
	return func(i int) int {
		return baseNumber * i
	}
}
