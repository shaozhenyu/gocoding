package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	s := Constructor([][]int{[]int{2, 0, 3, 0}})
	for i := 0; i < 5; i++ {
		fmt.Println(s.Pick())
	}
}

type Solution struct {
	sum   int
	s     []int
	rects [][]int
}

func Constructor(rects [][]int) Solution {
	sum := 0
	s := make([]int, len(rects))
	for i := 0; i < len(rects); i++ {
		r := rects[i]
		sum += (r[2] - r[0] + 1) * (r[3] - r[1] + 1)
		s[i] = sum
	}
	return Solution{
		rects: rects,
		s:     s,
		sum:   sum,
	}
}

func (this *Solution) Pick() []int {
	rand.Seed(time.Now().UnixNano())
	sum := rand.Intn(this.sum)
	var r []int
	for i := 0; i < len(this.s); i++ {
		if sum <= this.s[i] {
			r = this.rects[i]
			break
		}
	}
	x := r[0] + rand.Intn(r[2]-r[0]+1)
	y := r[1] + rand.Intn(r[3]-r[1]+1)
	return []int{x, y}
}
