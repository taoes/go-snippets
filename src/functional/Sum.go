package main

// 案例说明:使用闭包实现1~N 的累加
func main() {

	base := sum(1)

	for a := 0; a < 10; a++ {
		var s int
		s, _ = base(a)
		println(s)
	}

}

type Summer func(int) (int, Summer)

func sum(base int) Summer {
	return func(i2 int) (i int, summer Summer) {
		return i2 + base, sum(i + i2)
	}
}
