package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1,2,3,4,5,6}
	k := 4
	x := 5
	fmt.Println(findClosestElements(arr, k, x))
}

func findClosestElements(arr []int, k int, x int) []int {
	m := make(map[int][]int)
	closest := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		v := abs(x - arr[i])
		if _, ok := m[v]; !ok {
			m[v] = make([]int, 0)
			closest = append(closest, v)
		}
		m[v] = append(m[v], arr[i])
	}
	ret := make([]int, 0)
	sort.Ints(closest)
	for i := 0; i < len(closest); i++ {
		if len(ret) + len(m[closest[i]]) <= k {
			ret = append(ret, m[closest[i]]...)
		} else {
			t := m[closest[i]]
			sort.Ints(t)
			limit := k - len(ret)
			ret = append(ret, t[:limit]...)
			break
		}
	}
	sort.Ints(ret)
	return ret
}

func abs(x int) int {
	if x < 0 {
		x = -1 * x
	}
	return x
}
