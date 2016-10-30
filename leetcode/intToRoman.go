package main

import "fmt"

func main() {
	ret := intToRoman(2332)
	fmt.Println(ret)
}

func intToRoman(num int) string {

	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	str := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	s := ""

	if (len(nums) != len(str))  || num <= 0 {
		return ""
	}

	for num > 0 {
		for i := 0; i < len(nums); i++ {
			if num >= nums[i] {
				s += str[i]
				num -= nums[i]
				break
			}
		}
	}

	return s
}
