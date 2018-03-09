package main

import (
	"container/heap"
	"fmt"
)

func main() {
	grid := [][]int{{0,1,2,3,4},{24,23,22,21,5},{12,13,14,15,16},{11,17,18,19,20},{10,9,8,7,6}}
	//grid := [][]int{{10, 12, 4, 6}, {9, 11, 3, 5}, {1, 7, 13, 8}, {2, 0, 15, 14}}
	fmt.Println(swimInWater(grid))
}

type Item struct {
	x int
	y int
	v int
}

var direction = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].v < pq[j].v
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func swimInWater(grid [][]int) int {
	t := 2 << 31 * -1
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	start := &Item{
		x: 0,
		y: 0,
		v: grid[0][0],
	}
	heap.Push(&pq, start)

	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}
	for {
		this := heap.Pop(&pq).(*Item)
		if this.v > t {
			t = this.v
		}
		if this.x == m-1 && this.y == n-1 {
			return t
		}

		for _, d := range direction {
			x, y := this.x+d[0], this.y+d[1]
			if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] {
				visited[x][y] = true
				item := &Item{
					x: x,
					y: y,
					v: grid[x][y],
				}
				heap.Push(&pq, item)
			}
		}
	}
}
