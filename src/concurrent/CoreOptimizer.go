package main

import (
	"fmt"
	"runtime"
)

// GO 语言的多核优化特性
func main() {
	coreNumber := runtime.NumCPU()
	fmt.Printf("当前系统的核心:%v", coreNumber)
}
