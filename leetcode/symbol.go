package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fractionAddition("1/3+2/3-1/4"))
}

func fractionAddition(expression string) string {
	symbol := 1
	index := 0
	value := ""
	ch := make([]int, 0)
	mo := make([]int, 0)
	for i := 0; i < len(expression); i++ {
		b := expression[i]
		if b == '-' || b == '+' || b == '/' {
			if value != "" {
				v, _ := strconv.Atoi(value)
				if index%2 == 0 {
					ch = append(ch, v * symbol)
					symbol = 1
				} else {
					mo = append(mo, v)
				}
				index++
				value = ""
			}
			if b == '-' {
				symbol = -1
			} else if b == '+' {
				symbol = 1
			}
		} else {
			value += string(b)
		}
	}
	if value != "" {
		v, _ := strconv.Atoi(value)
		mo = append(mo, v)
	}
	sumMo := 1
	sumCh := 0
	for _, m := range mo {
		sumMo *= m
	}
	for i := 0; i < len(ch); i++ {
		sumCh += sumMo/mo[i] * ch[i]
	}
	d := division(sumMo, sumCh)
	return fmt.Sprintf("%d/%d", sumCh/d, sumMo/d)
}

func division(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
