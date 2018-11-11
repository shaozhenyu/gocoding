package main

import "fmt"

func main() {
	fmt.Println(countNumbersWithUniqueDigits(1))
	fmt.Println(countNumbersWithUniqueDigits(2))
	fmt.Println(countNumbersWithUniqueDigits(3))
	fmt.Println(countNumbersWithUniqueDigits(4))
}

// 10 10
// 91 9 * 9 + pre 第一位不能为10 排列组合
// 739 9 * 9 * 8 + pre
func countNumbersWithUniqueDigits(n int) int {
	if n == 1 {
		return 10
	}
	total := 10
	i := 2
	for i <= n {
		sum := 1
		for j := 0; j < i - 1; j++ {
			sum *= (9 - j)
		}
		sum *= 9
		total += sum
		i++
	}
	return total
}
