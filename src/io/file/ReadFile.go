package main

import (
	"io"
	"log"
	"os"
)

// 测试读取文件内容
func main() {
	file, e := os.Open("src/io/file/test/Test.txt")

	defer func() {
		log.Println("文件读取完成")
		file.Close()
	}();
	if e != nil {
		log.Printf("文件读取错误，请先写入文件到上述目录\n")
		log.Printf("错误信息:%v", e)
		return
	}

	bs := make([]byte, 1024, 1024)
	log.Println("============读取文件内容如下============")
	for {
		n, err := file.Read(bs)
		if n == 0 || err == io.EOF {
			return
		}
		log.Println(string(bs[:n]))
	}
}
