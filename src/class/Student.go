package main

import (
	"github.com/astaxie/beego/logs"
	"strings"
)

type Book struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

// 结构体的方法
func (b Book) desc() {
	logs.Info("id = ", b.Id, " Name = ", b.Name)
}

func (b Book) getUpperCaseName() string {
	return strings.ToUpper(b.Name)
}

func main() {

	b := Book{Id: 123, Name: "数据结构与算法经典问题解析"}
	logs.Info("the object's info = ")
	b.desc()
	logs.Info("the book's upper name = ", b.getUpperCaseName())

}
