package main

import "fmt"

func main() {
	n := 5
	mimes := [][]int{[]int{4, 2}}
	fmt.Println(orderOfLargestPlusSign(n, mimes))
}

func orderOfLargestPlusSign(N int, mines [][]int) int {
	zero := make(map[string]bool)
	for _, m := range mines {
		zero[fmt.Sprintf("%d+%d", m[0], m[1])] = true
	}
	var start, end, max int
	if N%2 != 0 {
		start = N / 2
	} else {
		start = N/2 - 1
	}
	end = N / 2
	max = start + 1
	now := 0
	hasDo := make(map[string]bool)
	for start >= 0 && end < N {
		for i := start; i <= end; i++ {
			for j := start; j <= end; j++ {
				if hasDo[fmt.Sprintf("%d+%d", i, j)] {
					continue
				}
				hasDo[fmt.Sprintf("%d+%d", i, j)] = true
				v := getK(i, j, N, zero)
				now = Max(now, v)
				if now >= max {
					return now
				}
			}
		}
		start--
		end++
		max -= 1
	}
	return now
}

func getK(i, j, N int, zero map[string]bool) int {
	for checkOne(i, j, N, zero) {
		up := next(i-1, j, -1, 0, N, zero)
		down := next(i+1, j, 1, 0, N, zero)
		x := min(up, down)
		left := next(i, j-1, 0, -1, N, zero)
		right := next(i, j+1, 0, 1, N, zero)
		y := min(left, right)
		return min(x, y) + 1
	}
	return 0
}

func next(i, j, x, y, N int, zero map[string]bool) (count int) {
	for checkOne(i, j, N, zero) {
		count++
		i += x
		j += y
	}
	return
}

func checkOne(i, j, N int, zero map[string]bool) bool {
	if !(i >= 0 && i < N) {
		return false
	}
	if !(j >= 0 && j < N) {
		return false
	}
	if zero[fmt.Sprintf("%d+%d", i, j)] {
		return false
	}
	return true
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
