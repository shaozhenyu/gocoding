package main

import "fmt"

func main() {
	r := []int{1,1,3,6,7,10,7,1,8,5,9,1,4,4,3}
	fmt.Println(rob(r))
}

func rob(nums []int) int {
	fmt.Println(nums)
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	step := make([][]int, len(nums))
	step1 := make([]int, len(nums))
	step1[0] = 0
	step1[1] = nums[1]
	step[0] = []int{nums[0], 1}
	if nums[1] >= nums[0] {
		step[1] = []int{nums[1], 0}
	} else {
		step[1] = []int{nums[0], 1}
	}
	i := 2
	for ; i < len(nums)-1; i++ {
		if step[i-1][0] > (step[i-2][0] + nums[i]) {
			step[i] = step[i-1]
		} else if step[i-1][0] < (step[i-2][0] + nums[i]){
			step[i] = []int{step[i-2][0] + nums[i], step[i-2][1]}
		} else {
			if step[i-1][1] == 0 {
				step[i] = step[i-1]
			} else {
				step[i] = []int{step[i-2][0] + nums[i], step[i-2][1]}
			}
		}
		step1[i] = max(step1[i-1], step1[i-2] + nums[i])
	}
	if len(nums) > 2 {
		step[i] = make([]int, 2)
		if step[i-2][1] == 0 {
			step[i-2][0] += nums[i]
		}
		step[i][0] = max(step[i-1][0], step[i-2][0])
		step1[i] = max(step1[i-1], step1[i-2] + nums[i])
	}
	return max(step[len(step)-1][0], step1[len(step1)-1])
}

func rob1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	step := make([]int, len(nums))
	step[0] = nums[0]
	step[1] = max(nums[1], nums[0])
	for i := 2; i < len(nums); i++ {
		step[i] = max(step[i-1], step[i-2] + nums[i])
	}
	return step[len(step)-1]
}

func rob(root *TreeNode) int {
	var leftC, rightC int
	var left, right int
	if root.Left != nil {
		left = root.Left.Val
		leftC = rob(root.Left)
	}
	if root.Right != nil {
		right = root.Right.Val
		rightC = rob(root.Right)
	}
	return max(root.Val + leftC + rightC, left + right)
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
