package main

import "fmt"

func main() {
	n := 3
	flights := [][]int{[]int{0, 1, 100}, []int{1, 2, 100}, []int{0, 2, 500}}
	src := 0
	dst := 2
	K := 1

	//	n = 5
	//	flights = [][]int{[]int{1,2,10},[]int{2,0,7},[]int{1,3,8},[]int{4,0,10},[]int{3,4,2},[]int{4,2,10},[]int{0,3,3},[]int{3,1,6},[]int{2,4,5}}
	//	src = 0
	//	dst = 4
	//	K =1
	fmt.Println(findCheapestPrice(n, flights, src, dst, K))
}

// dp
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	dp := make([][]int, K+2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 1 << 31
		}
	}
	dp[0][src] = 0
	for i := 1; i <= K+1; i++ {
		dp[i][src] = 0
		for j := 0; j < len(flights); j++ {
			f := flights[j]
			p := dp[i-1][f[0]] + f[2]
			if dp[i][f[1]] > p {
				dp[i][f[1]] = p
			}
		}
	}
	if dp[K+1][dst] == 1 << 31 {
		dp[K+1][dst] = -1
	}
	return dp[K+1][dst]
}

func findCheapestPrice1(n int, flights [][]int, src int, dst int, K int) int {
	graph := make([][]int, n)
	total := make([][]int, n)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		total[i] = make([]int, n)
		visited[i] = make([]bool, n)
	}
	for i := 0; i < len(flights); i++ {
		f := flights[i]
		graph[f[0]][f[1]] = f[2]
	}
	dfs(src, src, dst, K, 0, graph, total, visited)
	if total[src][dst] == 0 {
		total[src][dst] = -1
	}
	return total[src][dst]
}

func dfs(src, now, dst, k, price int, graph [][]int, total [][]int,  visited [][]bool) {
	if now == dst {
		if (price < total[src][dst]) || total[src][dst] == 0 {
			total[src][dst] = price
		}
		return
	}
	if k < 0 {
		return
	}
	for i := 0; i < len(graph[now]); i++ {
		if graph[now][i] > 0 {
			if visited[now][i] && ((price + graph[now][i] > total[now][i]) && total[now][i] > 0) {
				continue
			}
			visited[now][i] = true
			if (price + graph[now][i] < total[now][i]) || total[now][i] == 0 {
				total[now][i] = price + graph[now][i]
			}
			dfs(src, i, dst, k-1, price + graph[now][i], graph, total, visited)
		}
	}
}
