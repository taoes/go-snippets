package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("47.98.143.247")
	if ip == nil {
		println("invalid ip")
	} else {
		fmt.Printf("The add is %v ", ip.String())
	}
}
