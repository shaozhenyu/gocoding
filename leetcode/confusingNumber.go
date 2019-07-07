package main

import "fmt"

func main() {
	fmt.Println(confusingNumber(8111))
}

func confusingNumber(N int) bool {
	okm := map[int]int{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}
	old := N
	newOne := 0
	for N > 0 {
		t := N%10
		if v, ok := okm[t]; ok {
			newOne = newOne * 10 + v
		} else {
			return false
		}
		N = N/10
	}
	fmt.Println(newOne, old)
	return !(newOne == old)
}
