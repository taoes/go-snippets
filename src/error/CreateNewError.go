package main

import (
	"errors"
	"log"
)

func main() {

	number1, number2 := 3, 4
	result, err := Calculation(number1, number2)
	if err != nil {
		return
	}
	log.Printf("计算结果为:%v/%v=%v", number1, number2, result)

	_, err1 := Calculation(5, 0)
	if err1 != nil {
		log.Printf("程序出现异常，已退出", err1)
		return
	}

}

func Calculation(num1 int, num2 int) (result float32, err error) {
	if num2 == 0 {
		return 0, errors.New("除数不能为0")
	}
	return float32(num1) / float32(num2), nil
}
