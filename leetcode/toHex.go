package main

import "fmt"

func main() {
	fmt.Println(toHex(-123231))
	fmt.Println(toHex(-1))
}

func toHex(num int) string {
	if num < 0 {
		num = 1 << 32 + num
	}
	if num == 0 {
		return "0"
	}
	arr := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	str := ""
	for num != 0 {
		r := num%16
		str = arr[r] + str
		num = num/16
	}
	return str
}
