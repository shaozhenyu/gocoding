// 图 邻接矩阵 dfs
package main

import "fmt"

func main() {
	g := [][]int{[]int{0, 5}, []int{2, 4}, []int{2, 3}, []int{1, 2}, []int{0, 1}, []int{3, 4}, []int{3, 5}, []int{0, 2}}
	adj := Graph(6, g)
	marked := make([]bool, 6)
	edgeTo := make([]int, 6)
	dfs(0, adj, marked, edgeTo)
	fmt.Println(hasPathTo(5, marked))
	fmt.Println(pathTo(5, 0, edgeTo))
}

// v 顶点数 g 点到点的连接
func Graph(v int, G [][]int) [][]int {
	adj := make([][]int, v)
	for i := 0; i < v; i++ {
		adj[i] = make([]int, 0)
	}
	for _, g := range G {
		adj[g[0]] = append([]int{g[1]}, adj[g[0]]...)
		adj[g[1]] = append([]int{g[0]}, adj[g[1]]...)
	}
	return adj
}

func dfs(v int, adj [][]int, marked []bool, edgeTo []int) {
	marked[v] = true
	for _, w := range adj[v] {
		if !marked[w] {
			edgeTo[w] = v
			dfs(w, adj, marked, edgeTo)
		}
	}
}

func hasPathTo(v int, marked []bool) bool {
	return marked[v]
}

// e end, s start
func pathTo(e, s int, edgeTo []int) string {
	fmt.Println(edgeTo)
	path := ""
	for n := e; n != s; n = edgeTo[n] {
		path = fmt.Sprintf("%d-%s", n, path)
	}
	path = fmt.Sprintf("%d-%s", s, path)
	return path[:len(path)-1]
}
