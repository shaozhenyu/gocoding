package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := Constructor(10, 200)
	for i := 0; i < 2000; i++ {
		fmt.Println(s.Flip())
	}
	fmt.Println(len(s.test))
}

type Solution struct {
	n_rows, n_cols int
	sum            int
	black          map[int]int
	test           map[int]int
}

func Constructor(n_rows int, n_cols int) Solution {
	return Solution{
		n_rows: n_rows,
		n_cols: n_cols,
		sum:    n_rows * n_cols,
		black:  make(map[int]int),
		test:   make(map[int]int),
	}
}

func (this *Solution) Flip() []int {
	s := rand.Intn(this.sum)
	// fmt.Println("rand: ", s)
	tmp := s
	for {
		if v, ok := this.black[s]; ok {
			s = v
		} else {
			break
		}
	}
	if tmp <= this.sum-1 {
		// delete(this.black, this.sum)
		this.black[tmp] = this.sum - 1
	}
	this.test[s] = 1
	this.sum--
	// fmt.Println("s:", s, this.black, this.sum)
	return []int{s / this.n_cols, s % this.n_cols}
}

func (this *Solution) Reset() {
	this.sum = this.n_rows * this.n_cols
	this.black = make(map[int]int)
}
