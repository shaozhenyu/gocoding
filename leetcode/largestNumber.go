package main

import (
	"fmt"
	"sort"
	"strings"
	"strconv"
)

func main() {
	nums := []int{9, 91, 92, 991, 992, 0}
	fmt.Println(largestNumber(nums))
}

type Node struct {
	v string
	s string
}

func largestNumber(nums []int) string {
	ret := ""
	node := make([]Node, len(nums))
	for i := 0; i < len(nums); i++ {
		v := strconv.Itoa(nums[i])
		s := v + strings.Repeat(string(v[len(v)-1]), 10)
		node[i] = Node{v, s[:10]}
	}
	sort.Slice(node, func(i, j int) bool {
		return node[i].s > node[j].s
	})
	fmt.Println(node)
	for i := 0; i < len(node); i++ {
		ret += node[i].v
	}
	return ret
}
