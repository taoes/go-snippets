package main

import "sync"

var (
	num      int
	rwCounter sync.RWMutex
)

func main() {

	// 可以安全的设置
	setNumber(13)

	i := getNumber()
	print(i)
}

func getNumber2() int {
	rwCounter.RLock()
	defer rwCounter.RUnlock()
	return number
}

func setNumber2(num int) {
	rwCounter.RLock()
	number = num
	rwCounter.RUnlock()
}
