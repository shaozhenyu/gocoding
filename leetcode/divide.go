package main

import "fmt"

func main() {
	a :=213432413
	b := 33
	a = -2147483648
	b = -1
	fmt.Println(divide(a, b))
	fmt.Println(a/b)
}

func divide(dividend int, divisor int) int {
	flag := 0
	if dividend < 0 {
		dividend = 0 - dividend
		flag++
	}
	if divisor < 0 {
		divisor = 0 - divisor
		flag++
	}
	ret := 0
	ret = dd(dividend, divisor)
	if flag == -1 {
		ret = 0 - ret
	}
	if ret > (1 << 31) - 1 {
		ret = (1 << 31) - 1
	}
	return ret
}

func dd(dividend int, divisor int) int {
	if dividend < divisor {
		return 0
	}
	if dividend == divisor {
		return 1
	}
	ret := 0
	count := 1
	divisor1 := divisor
	for dividend > divisor1 {
		ret += count
		dividend -= divisor1
		divisor1 += divisor1
		count += count
	}
	return ret + dd(dividend, divisor)
}
