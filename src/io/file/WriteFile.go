package main

import (
	"fmt"
	"log"
	"os"
)

// 使用基本字节写入数据到文件

func main() {

	fileName := "src/io/file/test/九九乘法表.txt"

	_, e := os.Stat(fileName)
	if e == nil {
		// 文件存在尝试删除
		os.Remove(fileName)
		log.Println("文件删除完成")
	}

	file, error := os.Create("src/io/file/test/九九乘法表.txt")

	if error != nil {
		log.Println("文件创建失败")
		return
	}

	defer func() {
		file.Close()
		log.Println("文件写入完成")
	}()

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i >= j {
				s := fmt.Sprintf("%2v * %2v = %2v\t", i, j, i*j)
				file.WriteString(s)
			}
		}
		file.WriteString("\n")
	}

}
