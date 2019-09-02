package main

import "fmt"

func main() {
	i := printMessage(1, 2, 3)();
	print(i)
}

func printMessage(list ...int) func() int {
	for i := range list {
		fmt.Printf("接收到的参数：%v", i)
	}

	return func() int {
		return 3;
	}
}
