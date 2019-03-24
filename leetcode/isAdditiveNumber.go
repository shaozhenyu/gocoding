package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isAdditiveNumber("11235813"))
	fmt.Println(isAdditiveNumber("1023"))
}

func isAdditiveNumber(num string) bool {
	var checkNext func(x, y int64, num string) bool
	checkNext = func(x, y int64, num string) bool {
		if len(num) == 0 {
			return true
		}
		next := x + y
		if next < x || next < y {
			return false
		}
		nextStr := strconv.FormatInt(next, 10)
		if len(num) < len(nextStr) {
			return false
		}
		if nextStr == num[:len(nextStr)] {
			return checkNext(y, next, num[len(nextStr):])
		}
		return false
	}

	var x, y int64
	for i := 1; i < len(num); i++ {
		if num[0] == '0' && len(num[:i]) > 1 {
			break
		}
		x, _ = strconv.ParseInt(num[:i], 10, 64)
		for j := i+1; j < len(num); j++ {
			if num[i] == '0' && len(num[i:j]) > 1 {
				break
			}
			y, _ = strconv.ParseInt(num[i:j], 10, 64)
			ok := checkNext(x, y, num[j:])
			if ok {
				return true
			}
		}
	}
	return false
}
