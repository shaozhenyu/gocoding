package main

import (
	"fmt"
	"strconv"
	"math"
)

func main() {
	n := "1000000000000000000"
	n = "2251799813685247"
	n = "43"
	fmt.Println(smallestGoodBase(n))
}

func smallestGoodBase(n string) string {
	num, _ := strconv.ParseInt(n, 10, 64)
	fmt.Printf("%b\n", num)
	for i := 63; i >= 1; i-- {
		if k, ok := checkN(num, i); ok {
			return fmt.Sprintf("%d", k)
		}
		// break
	}
	return fmt.Sprintf("%d", num-1)
}

func checkN(num int64, size int) (int64, bool) {
	var left int64 = 2
	// var right int64 = num - 1
	var right int64 = int64(math.Pow(float64(num), math.Pow(float64(size - 1), -1))) + 1
	for left <= right {
		mid := left + (right - left)/2
		var now int64 = 1
		var sum int64 = 0
		for i := 0; i < size; i++ {
			sum += now
			now *= mid
		}
		fmt.Println(sum)
		if sum < num {
			left = mid + 1
		} else if sum > num {
			right = mid - 1
		} else {
			return mid, true
		}
	}
	return 0, false
}

func smallestGoodBase1(n string) string {
	num, _ := strconv.ParseInt(n, 10, 64)
	var i int64
	for i = 2; i <= num; i++ {
		if check(num, i) {
			return fmt.Sprintf("%d", i)
		}
	}
	return ""
}

func check(num int64, k int64) bool {
	for num > 0 {
		t := num%k
		if t != 1 {
			return false
		}
		num /= k
	}
	return true
}
