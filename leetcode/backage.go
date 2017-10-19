package main

import "fmt"

func main() {
	b := [][]int{[]int{5, 12}, []int{4, 3}, []int{7, 10}, []int{2, 3}, []int{6, 6}}
	fmt.Println(backage(b, 15))

	b1 := [][]int{[]int{5, 12, 3}, []int{4, 3, 1}, []int{7, 10, 1}, []int{2, 3, 1}, []int{6, 6, 1}}
	fmt.Println(backage2(b1, 15))
}

func backage2(nums [][]int,  total int) int {
	dp := make([]int, total+1)
	amount := make([]int, len(nums))
	count := make([]int, total+1)
	for i := 0; i < len(nums); i++ {
		amount[i] = nums[i][2]
	}

	for i := 0; i < len(nums); i++ {
		count = make([]int, total+1)
		for j := nums[i][0]; j <= total; j++ {
			if count[j-nums[i][0]] < amount[i] {
				dp[j] = max(dp[j], dp[j-nums[i][0]] + nums[i][1])
				if dp[j] == (dp[j-nums[i][0]] + nums[i][1]) {
					count[j] = count[j-nums[i][0]] + 1
				}
			} else if dp[j] == 0 {
				dp[j] = dp[j-1]
				count[j] = count[j-1]
			}
		}
	}
	return dp[total]
}

func backage1(nums [][]int, total int) int {
	dp := make([]int, total)
	for i := 0; i < len(nums); i++ {
		for j := nums[i][0]; j < total; j++ {
			dp[j] = max(dp[j], dp[j-nums[i][0]] + nums[i][1])
		}
	}
	fmt.Println(dp)
	return dp[total-1]
}

func backage(nums [][]int, total int) int {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, total)
	}

	for i := nums[0][0]; i < total; i++ {
		dp[0][i] = nums[0][1]
	}

	for i := 1; i < len(nums); i++ {
		for j := nums[i][0]; j < total; j++ {
			dp[i][j] = max(dp[i-1][j], nums[i][1]+dp[i-1][j-nums[i][0]])
		}
	}
	fmt.Println(dp)
	return dp[len(nums)-1][total-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
