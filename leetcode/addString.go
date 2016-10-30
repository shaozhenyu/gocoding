package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "9999"
	b := "1199999"
	aa := 9999
	bb := 1199999
	s := addStrings(a, b)
	fmt.Println(aa + bb)
	fmt.Println(s)
}

func addStrings(num1 string, num2 string) string {
	var n1, n2 string
	if len(num1) >= len(num2) {
		n1 = num1
		n2 = num2
	} else {
		n1 = num2
		n2 = num1
	}

	i := len(n1) - 1
	j := len(n2) - 1

	str := ""
	add := 0
	sum := 0

	for i >= 0 {
		if j >= 0 {
			sum = int(n1[i]) + int(n2[j]) - 2*48 + add
		} else {
			sum = int(n1[i]) - 48 + add
		}
		sumStr := strconv.Itoa(sum)

		if len(sumStr) == 2 {
			str = sumStr[1:] + str
			add = 1
		} else {
			str = sumStr + str
			add = 0
		}

		i--
		j--
	}

	if add == 1 {
		str = "1" + str
	}

	fmt.Println(str)

	return ""
}
