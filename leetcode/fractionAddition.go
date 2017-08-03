package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fractionAddition("-1/2+1/2-3/5+5/8"))
}

func fractionAddition(expression string) string {
	lenE := len(expression)
	num := ""
	nums := []int64{}

	var denominator int64 = 1
	for i := 0; i < lenE; i++ {
		if expression[i] != '/' {
			if num != "" && (expression[i] == '+' || expression[i] == '-') {
				d, _ := strconv.ParseInt(num, 10, 64)
				nums = append(nums, d)
				num = string(expression[i])

				denominator *= d
			} else {
				num += string(expression[i])
			}
		} else {
			d, _ := strconv.ParseInt(num, 10, 64)
			nums = append(nums, d)
			num = ""
		}
	}
	d, _ := strconv.ParseInt(num, 10, 64)
	nums = append(nums, d)
	denominator *= d

	var molecule int64 = 0
	for i := 0; i < len(nums); i = i + 2 {
		molecule += nums[i] * (denominator / nums[i+1])
	}

	if molecule == 0 {
		return "0/1"
	}

	common := GreatestCommonDivisor(molecule, denominator)

	fm := molecule / common
	fd := denominator / common

	return strconv.FormatInt(fm, 10) + "/" + strconv.FormatInt(fd, 10)
}

func GreatestCommonDivisor(a, b int64) int64 {
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}

	if a > b {
		a, b = b, a
	}

	for {
		c := b % a
		if c == 0 {
			return a
		} else {
			a = c
			b = a
		}
	}
	return 0
}
