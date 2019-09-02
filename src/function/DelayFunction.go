package main

import "fmt"

// 延迟函数的定义
func main() {

	// 利用延迟函数来修改返回值
	fmt.Printf("修改之后的值为:%v", modifyReturnValue())

	print("执行完成\n")
}

func modifyReturnValue() (value int) {
	defer func() {
		value++
	}()
	return
}
