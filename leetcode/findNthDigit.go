package main

import "fmt"

func main() {
	//fmt.Println(findNthDigit(3))
	//fmt.Println(findNthDigit(95))
	//fmt.Println(findNthDigit(96))
	//fmt.Println(findNthDigit(1053))
	//fmt.Println(findNthDigit(123))
	fmt.Println(findNthDigit(345435))
}

func findNthDigit(n int) int {
	start := 9
	step := 10
	add := 0
	mod := 1
	for n - mod * start > 0 {
		add += start
		n -= start * mod
		start *= step
		mod++
	}
	t := n/mod + add
	if n%mod == 0 {
		return t%10
	}
	t += 1
	m := mod - n%mod
	for m > 0 {
		t = t/10
		m--
	}
	return t%10
}
