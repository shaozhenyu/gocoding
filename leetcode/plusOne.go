package main

import "fmt"

func main() {
	a := []int{9, 9}
	fmt.Println(plusOne(a))
}

func plusOne(digits []int) []int {
	length := len(digits)
	add := 1
	for i := length-1; i >= 0; i-- {
		if i == 0 {
			if (digits[i] + add) == 10 {
				digits[i] = 1
				digits[length] = 0
			} else {
				digits[i] += 1
			}
			break
		}
		if (digits[i] + add) == 10 {
			add = 1
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}
	return digits
}
