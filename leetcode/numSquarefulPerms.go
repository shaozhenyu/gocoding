package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(numSquarefulPerms([]int{2,2,7,2}))
	fmt.Println(numSquarefulPerms([]int{2,7,9}))
}

func numSquarefulPerms(A []int) int {
	used := make([]bool, len(A))
	sum := 0
	sort.Ints(A)
	for i := 0; i < len(A); i++ {
		if i > 0 && A[i] == A[i-1] {
			continue
		}
		used[i] = true
		backtrack(A, used, A[i], 1, &sum)
		used[i] = false
	}
	return sum
}

func backtrack(A []int, used []bool, now int, count int, sum *int) {
	if count == len(A) {
		*sum += 1
		return
	}
	for i := 0; i < len(A); i++ {
		if used[i] {
			continue
		}
		var flag bool
		for j := i - 1; j >= 0; j-- {
			if !used[j] {
				if	A[j] == A[i] {
					flag = true
				}
				break
			}
		}
		if flag {
			continue
		}
		if isSquare(now + A[i]) {
			used[i] = true
			backtrack(A, used, A[i], count + 1, sum)
			used[i] = false
		}
	}
}

func isSquare(x int) bool {
	s := math.Sqrt(float64(x))
	i := float64(int(s))
	return s == i
}
