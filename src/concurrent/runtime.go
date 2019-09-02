package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	// 创建缓冲区大小为4 的通道
	intChannel := make(chan int, 4)
	// 标识等待 两个 goroutine 执行完成
	wg.Add(2)
	// 并发执行
	go completePercent(intChannel)
	go showPercent(intChannel)
	// 等待线程执行完成
	wg.Wait()

}

func showPercent(intChannel chan int) {
	// 通知主线程执行完成
	defer wg.Done()
	// 使用range 接收
	for percent := range intChannel {
		if percent == -1 {
			return
		}

		fmt.Printf("当前任务执行进度:")
		for a := 0; a < percent; a++ {
			fmt.Printf(">")
		}
		fmt.Printf(" %v%%\n",percent)
	}

}

// 延时模拟计算进度
func completePercent(intChannel chan int) {
	// 通知主线程执行完成
	defer wg.Done()
	percent := 0
	for {
		if percent >= 100 {
			fmt.Printf("计算完成\n")
			intChannel <- -1
			break
		}
		// 随机耗时，模拟计算操作
		delayTime := rand.Int31n(10)
		time.Sleep(time.Duration(delayTime) * time.Millisecond)
		// 写入数据到通道中
		percent++
		intChannel <- percent
	}
}
