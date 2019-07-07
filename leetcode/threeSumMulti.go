package main

import "fmt"

func main() {
	// fmt.Println(threeSumMulti([]int{1,1,2,2,3,3,4,4,5,5}, 8))
	fmt.Println(threeSumMulti([]int{95,93,46,2,29,41,28,74,9,10}, 196))
}

func threeSumMulti(A []int, target int) int {
	count := 0
	mod := 1000000007
	num := [101]int{}
	for i := 0; i < len(A); i++ {
		num[A[i]]++
	}
	for i := 0; i < len(num); i++ {
		if num[i] == 0 {
			continue
		}
		for j := i; j < len(num); j++ {
			if num[j] == 0 {
				continue
			}
			k := target - i - j
			if k < j {
				break
			}
			if k > 100 || num[k] == 0 {
				continue
			}
			count = (count + cal(i, j, k, num)) % mod
		}
	}
	return count
}

func cal(i, j, k int, nums [101]int) int {
	if i == j && j == k {
		num := nums[i]
		return num * (num-1) * (num-2) / (1 * 2 * 3)
	}
	if i == j && j != k {
		return nums[i] * (nums[i] - 1) * nums[k] / 2
	}
	if i != j && j == k {
		return nums[i] * nums[j] * (nums[j] - 1) / 2
	}
	return nums[i] * nums[j] * nums[k]
}
