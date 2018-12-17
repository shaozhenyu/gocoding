package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	buildings := [][]int{[]int{2, 9, 10}, []int{3, 7, 15}, []int{5, 12, 12}, []int{15, 20, 10}, []int{19, 24, 8}}
	buildings = [][]int{[]int{0,5,7},[]int{5,10,7},[]int{5,10,12},[]int{10,15,7},[]int{15,20,7},[]int{15,20,12},[]int{20,25,7}}
	buildings = [][]int{[]int{0,2,3},[]int{2,5,3}}
	fmt.Println(getSkyline(buildings))
}

type Node struct {
	x, y int
	flag int // 入度 出度
}

func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return [][]int{}
	}
	x := make([]Node, 0)
	for i := 0; i < len(buildings); i++ {
		x = append(x, Node{buildings[i][0], buildings[i][2], 1})
		x = append(x, Node{buildings[i][1], buildings[i][2], -1})
	}
	sort.Slice(x, func(i, j int) bool {
		if x[i].x != x[j].x {
			return x[i].x < x[j].x
		} else if x[i].flag == 1 && x[j].flag == 1 {
			return x[i].y > x[j].y
		} else if x[i].flag == -1 && x[j].flag == -1 {
			return x[i].y < x[j].y
		}
		if x[i].flag == 1 {
			return true
		}
		return false
	})
//	pos := 1
//	for i := 1; i < len(x); i++ {
//		if x[i-1].x != x[i].x {
//			fmt.Println(pos)
//			x[pos] = x[i]
//			pos++
//		} else {
//			if x[i-1].flag != x[i].flag {
//				if x[i-1].y == x[i].y {
//					pos--
//				}
//			}
//		}
//	}
//	x = x[:pos]
	fmt.Println(x)
	l := list.New()
	max := 0
	ret := make([][]int, 0)
	for i := 0; i < len(x); i++ {
		if x[i].flag == -1 {
			for e := l.Back(); e != nil; e = e.Prev() {
				v := e.Value.(Node)
				if v.y == x[i].y {
					l.Remove(e)
					break
				}
			}
			if l.Len() == 0 {
				ret = append(ret, []int{x[i].x, 0})
				max = 0
			} else {
				e := l.Back()
				v := e.Value.(Node)
				if max > v.y {
					ret = append(ret, []int{x[i].x, v.y})
				}
				max = v.y
			}
		} else {
			if x[i].y > max {
				max = x[i].y
				ret = append(ret, []int{x[i].x, x[i].y})
				l.PushBack(x[i])
			} else {
				preLen := l.Len()
				for e := l.Front(); e != nil; e = e.Next() {
					v := e.Value.(Node)
					if x[i].y < v.y {
						l.InsertBefore(x[i], e)
						break
					}
				}
				if l.Len() == preLen {
					l.PushBack(x[i])
				}
			}
		}
	}
	return ret
}
