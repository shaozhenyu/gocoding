package main

import (
	"fmt"
	"sort"
	// "container/list"
)

func main() {
	quality := []int{3,1,10,10,1}
	wage := []int{4,8,2,2,7}
	K := 3
	fmt.Println(mincostToHireWorkers(quality, wage, K))
}

type Node struct {
	quality int
	wage    int
	wq      float64
	index   int
}

func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	q := make([]Node, len(quality))
	for i := 0; i < len(quality); i++ {
		q[i] = Node{quality[i], wage[i], float64(wage[i])/float64(quality[i]), 0}
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i].quality < q[j].quality
	})
	wq := make([]Node, len(quality))
	sumKQ := 0
	for i := 0; i < len(q); i++ {
		if i < K - 1 {
			sumKQ += q[i].quality
		}
		wq[i] = Node{q[i].quality, q[i].wage, q[i].wq, i}
	}
	fmt.Println(wq)
	fmt.Println(q)
	fmt.Println(sumKQ)
	sort.Slice(wq, func(i, j int) bool {
		return wq[i].wq < wq[j].wq
	})
	
	last := K - 2
	var minQ float64 =  2 << 31 - 1
	for i := len(wq)-1; i >= K - 1; i-- {
		if wq[i].index <= last {
			sumKQ -= q[wq[i].index].quality
			last++
			for q[last].quality == -1 {
				last++
			}
			sumKQ += q[last].quality
		}
		q[wq[i].index].quality = -1
		v := float64(wq[i].wage) + wq[i].wq * float64(sumKQ)
		if minQ > v {
			minQ = v
		}
	}
	return minQ
}
