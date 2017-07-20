package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//结果
//10 1 2 3
//20 0 2 2
//2 0 2 2
//1 1 3 4

// 第一个defer会先调用calc里面的函数
// 所以第一个执行的应该是calc("10", 1, 2)
// 然后 a = 0
// 第二个defer同样先调用calc("20", 0, 2)
// 接下来defer遵循先进后出 calc("2", 0, 2)
// calc("1", 1, 3)
func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

// 2 0 calc("20", 0, 2)
// 20 0 2 2
// print 20 0 2 2
// print 2 0 1 1
