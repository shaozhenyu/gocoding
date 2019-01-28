package main

import "fmt"

func main() {
	x := 100
	target := 100000000
	fmt.Println(leastOpsExpressTarget(x, target))
}

func leastOpsExpressTarget(x int, target int) int {
	if target == x {
		return 0
	}
	count := 0
	if target > x && target < 2 * x {
		return leastOpsExpressTarget(x, target - x) + 1
	}
	if x > target {
		return min(2 * (x-target), target * 2 - 1)
	}
	t := x
	for t * x < target {
		count += 1
		t *= x
	}
	count++
	if t * x == target {
		return count
	}
	if t * x - target < target {
		return min(count + leastOpsExpressTarget(x, target - t), count + 1 + leastOpsExpressTarget(x, t * x - target))
	}
	return count + leastOpsExpressTarget(x, target - t)
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
