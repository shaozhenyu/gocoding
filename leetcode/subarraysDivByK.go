package main

import "fmt"

func main() {
	A := []int{0, -2, -3, -2, -3}
	A = []int{1, -2, 2, -1}
	//A = []int{10,10,10,10,10}
	K := 2
	fmt.Println(subarraysDivByK(A, K))
}

// 1
// 3
// 2
// 1
func subarraysDivByK(A []int, K int) int {
	m := make(map[int]int)
	sum := 0
	ret := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		v := (sum%K + K)%K
		if v == 0 {
			ret++
		}
		ret += m[v]
		m[v]++
	}
	return ret
}
