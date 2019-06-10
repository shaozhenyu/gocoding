package main

import "fmt"

func main() {
	fmt.Println(robotSim([]int{4,-1,3}, [][]int{}))
	fmt.Println(robotSim([]int{4,-1,4,-2,4}, [][]int{[]int{2,4}}))
}

func robotSim(commands []int, obstacles [][]int) int {
	left, right, up, down := [2]int{1, 0}, [2]int{-1, 0}, [2]int{0, 1}, [2]int{0, -1}
	obm := make(map[[2]int]bool)
	for i := 0; i < len(obstacles); i++ {
		obm[[2]int{obstacles[i][0], obstacles[i][1]}] = true
	}
	x, y, max := 0, 0, 0
	now := up
	for i := 0; i < len(commands); i++ {
		if commands[i] == -2 {
			if now == left {
				now = up
			} else if now == up {
				now = right
			} else if now == right {
				now = down
			} else {
				now = left
			}
		} else if commands[i] == -1 {
			if now == left {
				now = down
			} else if now == down {
				now = right
			} else if now == right {
				now = up
			} else {
				now = left
			}
		} else {
			for j := 1; j <= commands[i]; j++ {
				x += now[0]
				y += now[1]
				if obm[[2]int{x, y}] {
					x -= now[0]
					y -= now[1]
					break
				}
			}
			max = Max(max, x*x + y*y)
		}
	}
	return max
}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
