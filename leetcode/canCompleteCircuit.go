package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	total := 0
	has := 0
	idx := 0
	for i := 0; i < len(gas); i++ {
		has += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if has < 0 {
			has = 0
			idx = i + 1
		}
	}
	if total < 0 {
		return -1
	}
	return idx
}
