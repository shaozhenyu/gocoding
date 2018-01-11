package main

import (
	"fmt"
)

func main() {
	fmt.Println(findKthNumber(3, 3, 5))
}

func findKthNumber(m int, n int, k int) int {
	low := 1
	high := m * n + 1

	for low < high {
		mid := (low + high)/2
		count := less(mid, m, n)
		fmt.Println("count:", count, "mid:", mid)
		if count >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

// 计算比num小的值的总数
func less(num, m, n int) int {
	count := 0
	for i := 1; i <= m; i++ {
		fmt.Println()
		count += min(num/i, n)
		fmt.Println(count)
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
