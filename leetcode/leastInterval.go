package main

import (
	"fmt"
	"sort"
)

func main() {
	tasks := []byte{'A', 'B', 'A', 'B', 'C', 'A', 'B', 'C'}
	fmt.Println(leastInterval(tasks, 2))
}

func leastInterval(tasks []byte, n int) int {
	arr := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		arr[tasks[i]-'A']++
	}

	sort.Ints(arr)
	same := 0
	i := 24
	for i >= 0 && arr[i] == arr[25] {
		same++
		i--
	}
	return max(len(tasks), (arr[25]-1)*n+arr[25]+same)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
