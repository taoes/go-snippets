package main

import (
	"encoding/json"
	"log"
)

type Student struct {
	Name string
	Age  int
	Add  Address
}

type Address struct {
	Provide string
	City    string
}

func main() {
	stu := Student{
		Name: "周涛",
		Age:  26,
		Add: Address{
			Provide: "安徽省",
			City:    "亳州市",
		},
	}

	// 手动序列化
	bytes, e := json.Marshal(stu)
	if e != nil {
		log.Printf("序列化错误")
		return
	}
	jsonStr := string(bytes)
	println("序列话内容为:", jsonStr)

}
