package main

import (
	"fmt"
	"time"
)

func main() {

	// 使用匿名函数创建
	go func() {
		times := 0

		for {
			times++
			fmt.Printf("%v\n", times)
			time.Sleep(time.Second)
		}

	}()
	// 输入数据，卡主主线程，防止上面的goroutine启动后立刻停止
	var input string
	fmt.Scanln(&input)
}
