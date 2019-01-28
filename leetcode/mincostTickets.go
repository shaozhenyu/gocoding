package main

import "fmt"

func main() {
	days := []int{1,4,6,7,8,20}
	costs := []int{2,7,15}
	fmt.Println(mincostTickets(days, costs))
}

func mincostTickets(days []int, costs []int) int {
	visited := make(map[int]int)
	return mt(0, 0, days, costs, visited)
}

func mt(idx int, start int, days []int, costs []int, visited map[int]int) int {
	for idx < len(days) && days[idx] < start {
		idx++
	}
	if idx == len(days) {
		return 0
	}
	if visited[idx] > 0 {
		return visited[idx]
	}
	min := 2 << 31
	v := mt(idx+1, days[idx]+1, days, costs, visited) + costs[0]
	if min > v {
		min = v
	}
	v = mt(idx+1, days[idx]+7, days, costs, visited) + costs[1]
	if min > v {
		min = v
	}
	v = mt(idx+1, days[idx]+30, days, costs, visited) + costs[2]
	if min > v {
		min = v
	}
	visited[idx] = min
	return min
}
