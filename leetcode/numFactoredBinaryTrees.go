package main

import "fmt"

func main() {
	A := []int{2,5,10,100}
	fmt.Println(numFactoredBinaryTrees(A))
}

func numFactoredBinaryTrees(A []int) int {
	exist := make(map[int]bool)
	for i := 0; i < len(A); i++ {
		exist[A[i]] = true
	}
	father := make(map[int]map[int]bool)
	for i := 0; i < len(A); i++ {
		for j := i; j < len(A); j++ {
			v := A[i] * A[j]
			if exist[v] {
				if _, ok := father[v]; !ok {
					father[v] = make(map[int]bool)
				}
				father[v][A[i]] = true
				father[v][A[j]] = true
			}
		}
	}
	ret := 0
	t := make(map[int]int)
	for f := range father {
		ret = (ret + fn(father, f, t))%1000000007
	}
	return (len(A) + ret)%1000000007
}

func fn(father map[int]map[int]bool, f int, t map[int]int) int {
	if t[f] > 0 {
		return t[f]
	}
	child, ok := father[f]
	if !ok {
		return 0
	}
	count := 0
	used := make(map[int]bool)
	for k := range child {
		if used[k] {
			continue
		}
		k1 := f/k
		c1 := fn(father, k, t)
		t[k] = c1
		c2 := fn(father, k1, t)
		t[k1] = c2
		xx := 1
		if k1 != k {
			xx = 2
		}
		count += xx * (1 + c1 * c2 + (c1 + c2))
		used[k] = true
		used[k1] = true
	}
	return count
}
