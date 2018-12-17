package main

import "fmt"

func main() {
	nums := []int{-2, 1, -1, -2, -2}
	fmt.Println(circularArrayLoop(nums))
}

func circularArrayLoop(nums []int) bool {
	visited := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		newVisited := make([]bool, len(nums))
		newVisited[i] = true
		if move(i, nums[i] > 0, nums, visited, newVisited) {
			return true
		}
	}
	return false
}

func move(idx int, flag bool, nums []int, visited, newVisited []bool) bool {
	var next int
	next = (idx + nums[idx] + len(nums))%len(nums)
	if next == idx {
		return false
	}
	if flag != (nums[next] > 0) {
		return false
	}
	if newVisited[next] {
		return true
	}
	if visited[next] {
		return false
	}
	visited[next] = true
	newVisited[next] = true
	return move(next, flag, nums, visited, newVisited)
}
