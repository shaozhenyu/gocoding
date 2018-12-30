package main

import "fmt"

func main() {
	A := []int{1,2,3,4,5,6,7,8,9}
	// A = []int{32,98,9,43,66,49,83,94,95}
	fmt.Println(largestComponentSize(A))
}

func largestComponentSize(A []int) int {
	graph := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		graph[i] = make([]int, 0, len(A))
	}
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if check(A[i], A[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}
	m := make(map[int]map[int]struct{})
	visited := make([]bool, len(A))
	for i := 0; i < len(A); i++ {
		if visited[i] && len(graph[i]) == 0 {
			continue
		}
		visited[i] = true
		if _, ok := m[i]; !ok {
			m[i] = make(map[int]struct{})
		}
		// m[i][i] = struct{}{}
		traverse(graph, visited, m, i, i)
	}
	fmt.Println(graph)
	fmt.Println(m)
	max := 0
	for _, v := range m {
		if len(v) > max {
			max = len(v)
		}
	}
	return max
}

func traverse(graph[][]int, visited []bool, m map[int]map[int]struct{}, key, idx int) {
	// if visited[idx] {
	//     return
	// }
	m[key][idx] = struct{}{}
	for i := 0; i < len(graph[idx]); i++ {
		if visited[graph[idx][i]] {
			continue
		}
		visited[graph[idx][i]] = true
		m[key][graph[idx][i]] = struct{}{}
		traverse(graph, visited, m, key, graph[idx][i])
	}
}


func largestComponentSize1(A []int) int {
	graph := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		graph[i] = make([]int, 0)
	}
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if check(A[i], A[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}
	fmt.Println(graph)

	max := 0
	for i := 0; i < len(A); i++ {
		visited := make([]bool, len(A))
		c := 1 + dfs(i, visited, graph)
		if c > max {
			max = c
		}
	}
	return max
}

func dfs(idx int, visited []bool,  graph [][]int) int {
	visited[idx] = true
	ret := 0
	for i := 0; i < len(graph[idx]); i++ {
		if !visited[graph[idx][i]] {
			// fmt.Println(graph[idx])
			ret += dfs(graph[idx][i], visited, graph) + 1
		}
	}
	return ret
}

func check(x, y int) bool {
	if x > y {
		x, y = y, x
	}
	if y%x == 0 && x > 1 {
		return true
	}
	t := x/2
	for i := 2; i <= t; i++ {
		if x%i == 0 && y%i == 0 {
			return true
		}
	}
	return false
}
