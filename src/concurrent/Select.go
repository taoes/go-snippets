package main

import (
	"fmt"
	"time"
)

// 多路复用器
func main() {

	timeTicker := time.NewTicker(50 * time.Millisecond)

	i := 0

	for {
		select {
		case <-timeTicker.C:
			i++
			if i == 20 {
				goto stopTag
			}
			fmt.Printf("第%v次接收到数据\n", i)
		}
	}

stopTag:
	timeTicker.Stop()
	fmt.Printf("程序执行完成\n")

}
