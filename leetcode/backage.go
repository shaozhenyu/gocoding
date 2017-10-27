package main

import "fmt"

func main() {
	b := [][]int{[]int{5, 12}, []int{4, 3}, []int{7, 10}, []int{2, 3}, []int{6, 6}}
	fmt.Println(backpack1(b, 15))

	b1 := [][]int{[]int{5, 8, 2}, []int{4, 3, 2}, []int{7, 10, 2}, []int{2, 3, 3}, []int{6, 6, 3}}
	fmt.Println(backpack3(b1, 15))
}

func backpack3(nums [][]int, total int) int {
	// [][]int{[]int{5, 8, 2}, []int{4, 3, 2}, []int{7, 10, 2}, []int{2, 3, 3}, []int{6, 6, 3}}
	// nums[i][0]表示物品重量， nums[i][1]表示物品价值, nums[i][2]表示最大数量
	// 创建一个一维数组表示背包的容量为i时的最大价值
	dp := make([]int, total+1)

	// 创建一个数组表示容量为i时，放了多少个物品n
	count := make([]int, total+1)

	for i := 0; i < len(nums); i++ {
		// 每次重置count。每次循环表示一个物品存放的个数,所以需要每次重置
		count = make([]int, total+1)
		for j := nums[i][0]; j <= total; j++ {
			// count[j-nums[i][0]]表示已经放了物品i的个数， 如果大于物品i的总个数，跳过
			if count[j-nums[i][0]] < nums[i][2] {
				// 背包问题的通用公式,具体思路可以参考上面代码
				dp[j] = max(dp[j], dp[j-nums[i][0]]+nums[i][1])

				// 如果放入物品i价值更大，则count[j]需要在count[j-nums[i][0]](之前放入物品的格式)的基础上加1
				if dp[j] == (dp[j-nums[i][0]] + nums[i][1]) {
					count[j] = count[j-nums[i][0]] + 1
				}
			}

			if dp[j] == 0 || dp[j] < dp[j-1] {
				// 使用dp[i-1]的值
				dp[j] = dp[j-1]
				count[j] = count[j-1]
			}
		}
	}
	return dp[total]
}

func backpack1(nums [][]int, total int) int {
	// 创建一个一维数组，因为每个物品都有无数个，所以不需要二维数组，与本身比较即可。其余思路和前面的一样
	dp := make([]int, total+1)

	for i := 0; i < len(nums); i++ {
		for j := nums[i][0]; j <= total; j++ {
			dp[j] = max(dp[j], dp[j-nums[i][0]]+nums[i][1])
		}
		fmt.Println(dp)
	}
	return dp[total]
}

func backpack(nums [][]int, total int) int {
	// [][]int{[]int{5, 12}, []int{4, 3}, []int{7, 10}, []int{2, 3}, []int{6, 6}}  nums[i][0]表示物品重量， nums[i][1]表示物品价值
	// 有多少物品就有多少行
	dp := make([][]int, len(nums))

	// 列号代表背包的容量
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, total+1)
	}

	// 第一行放第一个物品，由于里面还没有物品，从能放下这个物品的重量开始，所以最优价值都是第一个物品本身的价值
	for i := nums[0][0]; i < total; i++ {
		dp[0][i] = nums[0][1]
	}

	// 遍历每一个物品
	for i := 1; i < len(nums); i++ {
		// 计算当背包容量为j时的最大值
		for j := nums[i][0]; j <= total; j++ {
			// dp[i-1][j]表示没有放这个物品时的最大价值， nums[i][1]+dp[i-1][j-nums[i][0]]: 该物品的价值+（背包重量-该物品的重量时的最大价值)
			// 判断将该物品放入背包还是不放入背包，取较大值
			dp[i][j] = max(dp[i-1][j], nums[i][1]+dp[i-1][j-nums[i][0]])
		}
	}
	return dp[len(nums)-1][total]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
