package main

import (
	"fmt"
	"sort"
)

func main() {

}

type Node struct {
	val int
	idx int
}

func validSubarrays(nums []int) int {
	node := make([]Node, len(nums)
	for i := 0; i < len(nums); i++ {
		node[i] := Node{nums[i], idx}
	}
	sort.Ints(node, func(i, j int) bool {
		if node[i].val == node[j].val {
			return node[i].idx < node[j].idx
		}
		return node[i].val < node[j].val
	})
	count := len(nums)

}
