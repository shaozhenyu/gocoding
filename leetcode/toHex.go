package main

import "fmt"

func main() {
	fmt.Println(toHex(-123231))
}

func toHex(num int) string {
	arr := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	str := ""
	for num != 0 {
		r := num%16
		str = arr[r] + str
		num = num/16
	}
	fmt.Println(str)
	return str
}
