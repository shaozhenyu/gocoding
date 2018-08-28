package main

import "fmt"

func main() {
	fmt.Println(allPathsSourceTarget([][]int{[]int{1,2}, []int{3}, []int{3}, []int{}}))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	s := make([]int, 20)
	return goPath(0, 0, s, graph, [][]int{})
}

func goPath(pos, index int, s []int, graph [][]int, ret [][]int) [][]int {
	if index == len(graph) - 1 {
		s[pos] = index
		tmp := make([]int, pos+1)
		copy(tmp, s)
		ret = append(ret, tmp)
		return ret
	}
	for _, key := range graph[index] {
		s[pos] = index
		ret = goPath(pos+1, key, s, graph, ret)
	}
	return ret
}
