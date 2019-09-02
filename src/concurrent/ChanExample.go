package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 主函数入口
func main() {

	// 定义存在缓冲区的通道，缓冲区内容为4
	buffChan := make(chan interface{}, 4)
	wg.Add(2)
	go buffChanSend(buffChan)
	go buffChanRec(buffChan)

	// 等到程序执行完成
	wg.Wait()
	fmt.Printf("程序执行完毕")
}

// 接受数据
func buffChanRec(rec chan interface{}) {
	defer wg.Done()

	for i := range rec {
		if i == nil {
			fmt.Printf("REC程序退出")
			break
		}
	}

}

// 发送数据
func buffChanSend(chantSend chan interface{}) {

	// 通知主线程执行完成的命令
	defer wg.Done()
	i := 1
	for {
		time.Sleep(time.Second)
		chantSend <- i
		fmt.Printf("发送数据:%v\n", i)
		if i > 7 {
			chantSend <- nil
			break
		}
		i = i + 1
	}

}
