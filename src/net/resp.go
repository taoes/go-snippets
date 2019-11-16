package main

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
)

/** 获取网络内容 */
func main() {

	resp, err := http.Get("https://api.unionfab.com")

	if err != nil {
		logs.Error("get web info fail, err info = ", err)
		return
	}

	body := resp.Body
	defer body.Close()
	respStr, _ := ioutil.ReadAll(resp.Body)
	logs.Info(string(respStr[:]))
}
