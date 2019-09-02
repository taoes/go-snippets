package main

import "flag"

func main() {

	host := flag.String("host", "", "")
	port := flag.Int("p", 8256, "this is an default port of proxy")

	flag.Parse()



	println(*port, *host)
}
