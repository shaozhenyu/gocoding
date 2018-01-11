package main

import "fmt"

func main() {
	n := 10
	//primes := []int{1, 2, 4, 7, 8, 13, 14, 16, 19, 26, 28, 32}
	primes := []int{2, 7, 13, 19}
	fmt.Println(nthSuperUglyNumber(n, primes))
}

func isUglyNumber(num int, m map[int]bool, primes []int) bool {
	if _, ok := m[num]; ok {
		return false
	}
	if num == 1 {
		return true
	}
	for i := 0; i < len(primes); i++ {
		if num%primes[i] == 0 {
			return isUglyNumber(num/primes[i], m, primes)
		}
	}
	return false
}

func nthSuperUglyNumber(n int, primes []int) int {
	ugly := make([]int, n)
	ugly[0] = 1

	bigger := make([]int, len(primes))

	for i := 1; i < n; i++ {
		ugly[i] = 2 << 32
		for j := 0; j < len(primes); j++ {
			ugly[i] = min(ugly[i], primes[j]*ugly[bigger[j]])
		}

		for j := 0; j < len(primes); j++ {
			for ugly[i] >= primes[j]*ugly[bigger[j]] {
				bigger[j]++
			}
		}
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
