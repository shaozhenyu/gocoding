package main

import (
	"fmt"
)

func main() {
	nums := []int{7,2,5,10,8,4,7}
	m := 4
	fmt.Println(splitArray(nums, m))
}

func splitArray(nums []int, m int) int {
	sum := 0
	sumArr := make([]int, len(nums))
	visited := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		visited[i] = make([]int, len(nums))
	}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sumArr[i] = sum
	}
	if m == 1 {
		return sum
	}
	min := sumArr[len(sumArr)-1]
	for i := 0; i < len(nums); i++ {
		v := split(i+1, nums, sumArr, m - 1, sumArr[i], visited)
		fmt.Println("aa:", v, i)
		if min > v {
			min = v
		}
	}
	return min
}

func split(start int, nums, sum []int, m int, max int, visited [][]int) int {
	fmt.Println(start, m, max)
	if start >= len(nums) {
		return max
	}
	if visited[m][start] > max {
		// if visited[m][start]
		fmt.Println("vvvvv: ", m, start, visited[m][start])
		return visited[m][start]
	}
	if m == 1 {
		v := sum[len(sum)-1]
		if start > 0 {
			v -= sum[start-1]
		}
		 // visited[m][start] = v
		if v > max {
			max = v
		}
		return max
	}
	t := sum[len(sum)-1]
	for i := start; i < len(nums)-1; i++ {
		now := sum[i]
		if start > 0 {
			now -= sum[start-1]
		}
		t1 := split(i+1, nums, sum, m-1, now, visited)
		if t1 < t {
			t = t1
		}
	}
	visited[m][start] = t
	if t > max {
		max = t
	}
	return max
}
