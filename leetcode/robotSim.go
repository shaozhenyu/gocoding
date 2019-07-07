package main

import "fmt"

func main() {

}

func robotSim(commands []int, obstacles [][]int) int {
	obs := make(map[[2]int]struct{}{})
	for i := 0; i < len(obstacles); i++ {
		obs[[2]int{obstacles[i][0]], obs[obstacles[i][1]}] = struct{}{}
	}
	start := [2]int{0, 0}
	dic := [2]int{1, 0}
	right := map[[2]int][2]int{
		[2]int{1, 0}: [2]int{0, 1},
		[2]int{0, 1}: [2]int{-1, 0},
		[2]int{-1, 0}: [2]int{0, -1},
		[2]int{0, -1}: [2]int{1, 0},
	}
	left := map[[2]int][2]int{
		[2]int{1, 0}: [2]int{0, -1},
		[2]int{0, -1}: [2]int{-1, 0},
		[2]int{-1, 0}: [2]int{0, 1},
		[2]int{0, 1}: [2]int{1, 0},
	}
	for _, c := range commands {
		if c == -1 {
			dic = right[dic]
		} else if c == -2 {
			dic = left[dic]
		} else {
			if dic[0] != 0 {
				for i := 1; i <= c; i++ {
					if obs[s:]
				}
			}
		}
	}
}
