package main

import "fmt"

func main() {
	m := []int{0, 0, 0, 0, 25}
	fmt.Println(findMinMoves(m))
}

func findMinMoves(machines []int) int {
	fmt.Println("origin: ", machines)
	sum := 0
	length := len(machines)
	for i := 0; i < len(machines); i++ {
		sum += machines[i]
	}
	if sum%length != 0 {
		return -1
	}
	avg := sum/length
	count := 0
	ret := 0
	for i := 0; i < len(machines); i++ {
		count += machines[i] - avg
		ret = max(ret, max(abs(count), machines[i] - avg))
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
