package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
)

// GoLang 高并发写入文件
func main() {
	chanInt := make(chan int)
	wg := sync.WaitGroup{}

	for a := 0; a < 100; a++ {
		wg.Add(1)
		go productData(chanInt, &wg)
	}
	go writeFile(chanInt)

	go func() {
		wg.Wait()
		close(chanInt)
	}()

}

func writeFile(chanInt chan int) {

	f, err := os.Create("src/io/file/test/concurrent.txt")
	defer f.Close()
	if err != nil {
		log.Printf("创建文件异常")
	}

	for d := range chanInt {
		fmt.Fprintln(f, d)
	}

}

func productData(chanInt chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	intn := rand.Intn(129)
	chanInt <- intn
}
