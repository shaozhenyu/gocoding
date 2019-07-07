package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	// s := Constructor(10, []int{1, 5, 6, 9})
	s := Constructor(3, []int{0})
	fmt.Println(s)
	for i := 0; i < 5; i++ {
		fmt.Println(s.Pick())
	}
}

type Solution struct {
	reBlack map[int]int
	total int
}


func Constructor(N int, blacklist []int) Solution {
	bm := make(map[int]int)
	for i := 0; i < len(blacklist); i++ {
		bm[blacklist[i]] = -1
	}
	total := N - len(bm)
	for i := 0; i < len(blacklist); i++ {
		if blacklist[i] >= total {
			continue
		}
		if blacklist[i] < total {
			for {
				if _, ok := bm[N-1]; ok {
					N--
				} else {
					break
				}
			}
			bm[blacklist[i]] = N-1
			N--
		}
	}
	return Solution{
		reBlack: bm,
		total: total,
	}
}


func (this *Solution) Pick() int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(this.total)
	if v, ok := this.reBlack[n]; ok{
		return v
	}
	return n
}
