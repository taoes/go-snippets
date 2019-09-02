package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

/**
等待组
*/

var urls = []string{
	"http://www.github.com/",
	"https://www.qiniu.com/",
	"https://www.golangtc.com/",
}

// 定义等待组
var wg sync.WaitGroup

func main() {

	for _, url := range urls {
		wg.Add(1)
		// 开启一个goroutine
		go func(url string) {
			// 计数器减1，等价于 wg.Add(-1)
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				fmt.Errorf("请求:%v数据异常\n", url)
			} else {
				fmt.Printf("请求完成,code=%v\n", resp.Status)
				defer resp.Body.Close()
				s, _ := ioutil.ReadAll(resp.Body)
				fmt.Print(string(s))
			}
		}(url)
	}
	// 阻塞WG的计数器到0
	wg.Wait()
	fmt.Printf("程序执行完成....")
}
