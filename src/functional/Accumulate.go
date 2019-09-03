package main

import (
	"log"
)

// 闭包的变量具有记忆效应
func main() {

	// 初始化闭包，返回一个累加器
	addFunc := add(10)
	for i := 0; i < 10; i++ {
		log.Printf("当前调用value = %v\n", addFunc())
	}

}

type AddFunc func() int

func add(number int) AddFunc {
	return func() int {
		number = number + 1
		return number
	}
}
