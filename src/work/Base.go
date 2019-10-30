package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

func main() {
	fmt.Print(aurora.Blue("正在读取远程服务器配置信息:\n"))

	for i := 0; i < 10; i++ {
		fmt.Print(aurora.Cyan(fmt.Sprintf("【%d】服务器:%d IP:%d \n", i, i, i)))
	}

	fmt.Print(aurora.Blue("请输入需要访问的服务器:【 】"))

}
