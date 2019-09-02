package main

import "sync"

var (
	number      int
	counterLock sync.Mutex
)

func main() {

	// 可以安全的设置
	setNumber(13)

	i := getNumber()
	print(i)
}

func getNumber() int {
	counterLock.Lock()
	defer counterLock.Unlock()
	return number
}

func setNumber(num int) {
	counterLock.Lock()
	number = num
	counterLock.Unlock()
}
