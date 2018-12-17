package main

import (
	"fmt"
	"math"
)


func main() {
	nums := []int{}
	na := Constructor(nums)
	fmt.Println(na.stree)
	return
	fmt.Println(na.SumRange(2, 3))
	na.Update(2, 100)
	fmt.Println(na.stree)
	fmt.Println(na.nums, na.SumRange(2, 3))
}

type NumArray struct {
	stree []int
	nums []int
}


func Constructor(nums []int) NumArray {
	if len(nums) <= 0 {
		return NumArray{}
	}
	high := math.Ceil(math.Log2(float64(len(nums))))
	size := 2 * int(math.Pow(2, high)) - 1
	stree := make([]int, size)
	build(stree, nums, 0, 0, len(nums)-1)
	return NumArray{stree, nums}
}

func build(t, nums []int, idx, left, right int) int {
	fmt.Println(idx, left)
	if left == right {
		t[idx] = nums[left]
		return t[idx]
	}
	mid := left + (right - left)/2
	sum := build(t, nums, idx << 1 + 1, left, mid) + build(t, nums, idx << 1 + 2, mid+1, right)
	t[idx] = sum
	return sum
}


func (this *NumArray) Update(i int, val int)  {
	old := this.nums[i]
	add := val - old
	this.nums[i] = val
	this.updateNode(0, i, 0, len(this.nums)-1, add)
}

func (this *NumArray) updateNode(sIdx, nIdx, left, right, add int) {
	if nIdx < left || nIdx > right {
		return
	}
	this.stree[sIdx] += add
	if left == right {
		return
	}
	mid := left + (right - left)/2
	this.updateNode(sIdx << 1 + 1, nIdx, left, mid, add)
	this.updateNode(sIdx << 1 + 2, nIdx, mid+1, right, add)
}


func (this *NumArray) SumRange(i int, j int) int {
	st := this.stree
	return getSum(st, i, j, 0, len(this.nums)-1, 0)
}

func getSum(st []int, i, j, start, end, idx int) int {
	if i > end || j < start {
		return 0
	}
	if i <= start && j >= end {
		return st[idx]
	}
	mid := start + (end - start)/2
	return getSum(st, i, j, start, mid, idx << 1 + 1) + getSum(st, i, j, mid+1, end, idx << 1 + 2)
}

