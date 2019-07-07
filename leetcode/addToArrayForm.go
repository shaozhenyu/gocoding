package main

import "fmt"

func main() {
	fmt.Println(addToArrayForm([]int{1,2,0,0}, 34))
}

func addToArrayForm(A []int, K int) []int {
	k := make([]int, 0, 10)
	for K > 0 {
		k = append(k, K%10)
		K /= 10
	}
	for i := 0; i < len(k)/2; i++ {
		k[i], k[len(k)-i-1] = k[len(k)-i-1], k[i]
	}
	ret := make([]int, 20000)
	i, j := len(A) - 1, len(k) - 1
	idx := 0
	add := 0
	for i >= 0 && j >= 0 {
		ret[idx] = A[i] + k[j] + add
		if ret[idx] >= 10 {
			add = 1
			ret[idx] -= 10
		} else {
			add = 0
		}
		i--
		j--
		idx++
	}
	for i >= 0 {
		ret[idx] = A[i] + add
		if ret[idx] >= 10 {
			add = 1
			ret[idx] -= 10
		} else {
			add = 0
		}
		i--
		idx++
	}
	for j >= 0 {
		ret[idx] = k[j] + add
		if ret[idx] >= 10 {
			add = 1
			ret[idx] -= 10
		} else {
			add = 0
		}
		j--
		idx++
	}
	if add == 1 {
		ret[idx] = 1
		idx++
	}
	ret = ret[:idx]
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret) - i - 1] = ret[len(ret) - i - 1], ret[i]
	}
	return ret
}
