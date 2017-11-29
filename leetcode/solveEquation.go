package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solveEquation("x+5-3+x=6+x-2"))
}

func solveEquation(equation string) string {
	v := strings.Split(equation, "=")
	left := v[0]
	right := v[1]

	leftX, leftD := cal(left)
	rightX, rightD := cal(right)

	x := leftX - rightX
	d := rightD - leftD

	if x == 0 && d == 0 {
		return "Infinite solutions"
	}
	if x == 0 {
		return "No solution"
	}
	if d == 0 {
		return "x=0"
	}
	return "x=" + strconv.Itoa(d/x)
}

func cal(nums string) (x, d int) {
	symbol := 1
	start := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == byte('+') || nums[i] == byte('-') || i == len(nums)-1 {
			pre := nums[start:i]
			if i == len(nums)-1 {
				pre = nums[start:]
			}
			if strings.Contains(pre, "x") {
				if pre[0] == byte('x') {
					x += 1 * symbol
				} else {
					preD, _ := strconv.Atoi(pre[:len(pre)-1])
					x += preD * symbol
				}
			} else {
				preD, _ := strconv.Atoi(pre)
				d += preD * symbol
			}

			if nums[i] == byte('+') {
				symbol = 1
			} else if nums[i] == byte('-') {
				symbol = -1
			}
			start = i + 1
		}
	}
	return
}
