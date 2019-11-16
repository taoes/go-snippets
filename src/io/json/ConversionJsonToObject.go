package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Address struct {
		Provide string
		City    string
	}

	type Student struct {
		Name string
		Age  int
		Add  Address
	}

	data := []byte(`{"Name":"周涛","Age":26,"Add":{"Provide":"安徽省","City":"亳州市"}}`)

	stu := Student{}
	_ = json.Unmarshal(data, &stu)
	fmt.Printf("反序列化结果为:%v", stu)
}
