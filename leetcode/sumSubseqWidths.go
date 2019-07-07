package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sumSubseqWidths([]int{1, 2, 3, 4, 5}))
	fmt.Println(sumSubseqWidths([]int{1, 2, 3, 4, 5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30}))
	fmt.Println(sumSubseqWidths([]int{1,5,6,9,11,23}))
}

func sumSubseqWidths(A []int) int {
	mod := 1000000007
	sort.Ints(A)
	ret, sum := 0, 0
	p := 1
	for i := len(A) - 2; i >= 0; i-- {
		p = (p * 2) % mod
		sum = (sum * 2 % mod + (p - 1) * (A[i+1] - A[i]) % mod) % mod
		ret = (ret + sum) % mod
	}
	return ret
}

func sumSubseqWidths2(A []int) int {
	mod := 1000000007
	sort.Ints(A)
	ret, sum := 0, 0
	p := 1
	for i := 1; i < len(A); i++ {
		sum = (sum + (A[i]-A[0])*p) % mod
		p = (p * 2) % mod
	}
	fmt.Println("xx: ", sum)
	ret = (ret + sum) % mod
	for i := 1; i < len(A); i++ {
		sum = (sum - (p-1)*(A[i]-A[i-1])) / 2
		fmt.Println("xx: ", sum)
		ret = (ret + sum) % mod
		p /= 2
	}
	return ret
}

func sumSubseqWidths1(A []int) int {
	mod := 1000000007
	sort.Ints(A)
	ret, sum := 0, 0
	p := 1
	for i := 1; i < len(A); i++ {
		sum = (sum + (A[i]-A[0])*p) % mod
		p = (p * 2) % mod
	}
	ret = (ret + sum) % mod
	// fmt.Println(sum)
	for i := 1; i < len(A); i++ {
		sum = (sum - (p-1)*(A[i]-A[i-1])) / 2
		ret = ret + sum
		// fmt.Println(sum)
	}
	return ret
}
