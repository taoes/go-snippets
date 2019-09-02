package main

/** 命名函数返回值 */
func main() {

	print(parse(123))
}

func parse(number int) (newNumber int) {
	print(number)
	newNumber = 111 + number
	return
}
