package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {

	// 初始化伪随机种子
	rand.Seed(time.Now().Unix())
	number := rand.Int31n(30)
	_, e := testError(number)
	if e != nil {
		log.Printf("执行程序发生异常", e)
	}
}

func testError(number int32) (bool, error) {

	if number%2 == 0 {
		log.Printf("触发异常")
		return false, &ErrorMessage{"", 3}
	}
	return true, nil
}

type ErrorMessage struct {
	message string
	line    int
}

func (errorMessage *ErrorMessage) Error() string {
	return fmt.Sprintf("发生了异常 %v:%v", errorMessage.message, errorMessage.line)
}
