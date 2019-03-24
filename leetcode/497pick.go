package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := Constructor([][]int{[]int{1, 1, 5, 5}})
	s = Constructor([][]int{[]int{99358434, 62418790, 99360410, 62419739}, []int{9949520, 63556732, 9949788, 63556965}})
	s = Constructor([][]int{[]int{-2, -2, -1, -1}, []int{1, 0, 3, 0}})
	fmt.Println(s.Pick())
}

type Solution struct {
	s [][]int
}

func Constructor(rects [][]int) Solution {
	return Solution{
		s: rects,
	}
}

func (this *Solution) Pick() []int {
	rand.Seed(time.Now().UnixNano())
	r := this.s[rand.Intn(len(this.s))]
	x := r[0]
	if r[2] != r[0] {
		x = r[0] + rand.Intn(r[2] - r[0])
	}
	y := r[1]
	if r[1] != r[3] {
		y = r[1] + rand.Intn(r[3] - r[1])
	}
	fmt.Println(x, y)
	return []int{x, y}
}
