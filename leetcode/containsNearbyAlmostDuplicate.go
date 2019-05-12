package main

import "fmt"

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1,2,1,0}, 1, 1))
}

type TreeNode struct {
	idx map[int]bool
	val int
	left *TreeNode
	right *TreeNode
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) == 0 {
		return false
	}
	root := &TreeNode{map[int]bool{0: true}, nums[0], nil, nil}
	for i := 1; i < len(nums); i++ {
		root.insert(i, nums[i])
	}
	for i := 0; i < len(nums); i++ {
		if root.check(i, nums[i], k, t) {
			return true
		}
	}
	return false
}

func (root *TreeNode) check(idx, val, k, t int) bool {
	if root == nil {
		return false
	}
	if abs(root.val - val) <= t {
		for idx1 := range root.idx {
			if idx1 != idx && abs(idx1 - idx) <= k {
				return true
			}
		}
	}
	if root.val < val {
		if root.right.check(idx, val, k, t) {
			return true
		}
		if val - root.val <= t {
			if root.left.check(idx, val, k, t) {
				return true
			}
		}
	}
	if root.val > val {
		if root.left.check(idx, val, k, t) {
			return true
		}
		if root.val - val <= t {
			if root.right.check(idx, val, k, t) {
				return true
			}
		}
	}
	return false
}

func (root *TreeNode) insert(idx, val int) {
	if root.val == val {
		root.idx[idx] = true
		return
	}
	if root.val > val {
		if root.left == nil {
			root.left = &TreeNode{map[int]bool{idx: true}, val, nil, nil}
		} else {
			root.left.insert(idx, val)
		}
	}
	if root.val < val {
		if root.right == nil {
			root.right = &TreeNode{map[int]bool{idx: true}, val, nil, nil}
		} else {
			root.right.insert(idx, val)
		}
	}
}

func (root *TreeNode) print() {
	if root != nil {
		fmt.Println(root.val, root.idx)
		root.left.print()
		root.right.print()
	}
}


func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		j := i + 1
		if i > 0 && nums[i] == nums[i-1] {
			j = i + k - 1
		}
		for ; j <= i + k && j < len(nums); j++ {
			if abs(nums[i] - nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
