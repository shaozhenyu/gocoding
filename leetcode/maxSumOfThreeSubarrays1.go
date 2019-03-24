package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 2, 6, 7, 5, 1}
	k := 2
	nums = []int{17, 7, 19, 11, 1, 19, 17, 6, 13, 18, 2, 7, 12, 16, 16, 18, 9, 3, 19, 5}
	k = 6
	fmt.Println(maxSumOfThreeSubarrays(nums, k))
}

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	sum := make([]int, len(nums)-k+1)
	max := make(map[int]map[int]Node)
	for i := 0; i < k; i++ {
		sum[0] += nums[i]
	}
	max[0] = make(map[int]Node)
	for i := k; i < len(nums); i++ {
		sum[i-k+1] = nums[i] + sum[i-k] - nums[i-k]
		max[i-k+1] = make(map[int]Node)
	}
	return maxSum3(k, sum, max)
}

type Node struct {
	sum int
	idx1, idx2, idx3 int
}

func maxSum3(k int, sum []int, max map[int]map[int]Node) []int {
	maxVal := -1 << 32
	var idx1, idx2, idx3 int
	for i := 0; i < len(sum); i++ {
		val2, i2, i3 := maxSum2(k, i, sum, max)
		if val2+sum[i] > maxVal {
			idx1, idx2, idx3 = i, i2, i3
			maxVal = val2 + sum[i]
		}
	}
	return []int{idx1, idx2, idx3}
}

func maxSum2(k int, idx int, sum []int, max map[int]map[int]Node) (int, int, int) {
	if v, ok := max[idx][2]; ok {
		return v.sum, v.idx2, v.idx3
	}
	maxVal := -1 << 32
	idx2, idx3 := 0, 0
	for i := idx + k; i < len(sum); i++ {
		val3, i3 := maxSum1(k, i, sum, max)
		if sum[i]+val3 > maxVal {
			idx2, idx3 = i, i3
			maxVal = sum[i] + val3
		}
	}
	max[idx][2] = Node{sum: maxVal, idx2: idx2, idx3: idx3}
	return maxVal, idx2, idx3
}

func maxSum1(k int, idx int, sum []int, max map[int]map[int]Node) (int, int) {
	if v, ok := max[idx][1]; ok {
		return v.sum, v.idx3
	}
	maxVal := -1 << 32
	maxIdx := 0
	for i := idx + k; i < len(sum); i++ {
		if sum[i] > maxVal {
			maxVal = sum[i]
			maxIdx = i
		}
	}
	max[idx][1] = Node{sum: maxVal, idx3: maxIdx}
	return maxVal, maxIdx
}
