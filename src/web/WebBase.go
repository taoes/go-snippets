package main

import (
	"fmt"
	"net/http"
)

func handleIndexRequest(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, "Hello,World")
}

func main() {
	http.HandleFunc("/", handleIndexRequest)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Errorf("服务器启动错误")
	} else {
		fmt.Print("Web 服务成功启动在9090")
	}
}
