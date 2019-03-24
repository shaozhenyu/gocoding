package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(carFleet(10, []int{2, 4}, []int{3, 2}))
	// fmt.Println(carFleet(31, []int{5,26,18,25,29,21,22,12,19,6}, []int{7,6,6,4,3,4,9,7,6,4}))
	// fmt.Println(carFleet(44, []int{10, 43, 34, 4, 9, 35, 29, 5, 3, 2, 1, 41, 38, 6}, []int{6, 5, 1, 1, 7, 1, 10, 3, 9, 8, 2, 9, 5, 10}))
	fmt.Println(carFleet(55680, []int{53519,40853,7008,47279,31279,36173,16808,2024,21908,21794}, []int{524396,467454,287575,112625,961858,134139,415285,182401,334920,70493}))
}

type Node struct {
	pos   int
	speed int
}

func carFleet(target int, position []int, speed []int) int {
	node := make([]Node, len(position))
	for i := 0; i < len(position); i++ {
		node[i] = Node{position[i], speed[i]}
	}
	sort.Slice(node, func(i, j int) bool {
		return node[i].pos > node[j].pos
	})
	count := 0
	for len(node) > 1 {
		pre := node[0]
		prePos := pre.pos + pre.speed
		pre.pos = prePos
		tmp := make([]Node, 0, len(node))
		for i := 1; i < len(node); i++ {
			now := node[i]
			nowPos := now.pos + now.speed
			now.pos = nowPos
			if nowPos >= prePos {
				if nowPos > target {
					if float64(target-(now.pos-now.speed))/float64(now.speed) >= float64(target-(pre.pos-pre.speed))/float64(pre.speed) {
						count++
						pre = now
						prePos = nowPos
					}
				}
			} else {
				if prePos >= target {
					count++
				} else {
					tmp = append(tmp, pre)
				}
				pre = now
				prePos = nowPos
			}
		}
		if prePos >= target {
			count++
		} else {
			tmp = append(tmp, pre)
		}
		node = tmp
	}
	if len(node) == 1 {
		count++
	}
	return count
}
