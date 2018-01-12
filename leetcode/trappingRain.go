package main

import (
	"container/heap"
	"fmt"
)

type cell struct {
	row    int
	col    int
	height int
}

func NewCell(row, col, height int) *cell {
	return &cell{
		row:    row,
		col:    col,
		height: height,
	}
}

func (c *cell) Print() {
	fmt.Println(c.row, c.col, c.height)
}

type PriorityQueue []*cell

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].height < p[j].height
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	item := x.(*cell)
	*p = append(*p, item)
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}

func main() {
	heights := [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}
	fmt.Println(trapRainWater(heights))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}

	m := len(heightMap)
	n := len(heightMap[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	p := make(PriorityQueue, m*2+n*2)

	var pos int
	for i := 0; i < m; i++ {
		visited[i][0] = true
		visited[i][n-1] = true
		p[pos] = NewCell(i, 0, heightMap[i][0])
		pos++
		p[pos] = NewCell(i, n-1, heightMap[i][n-1])
		pos++
		//heap.Push(&p, NewCell(i, 0, heightMap[i][0]))
		//heap.Push(&p, NewCell(i, n-1, heightMap[i][n-1]))
	}

	for i := 0; i < n; i++ {
		visited[0][i] = true
		visited[m-1][i] = true
		p[pos] = NewCell(0, i, heightMap[0][i])
		pos++
		p[pos] = NewCell(m-1, i, heightMap[m-1][i])
		pos++
		//heap.Push(&p, NewCell(0, i, heightMap[0][i]))
		//heap.Push(&p, NewCell(m-1, i, heightMap[m-1][i]))
	}

	heap.Init(&p)
	step := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	ret := 0
	for p.Len() > 0 {
		c := heap.Pop(&p).(*cell)
		for i := 0; i < len(step); i++ {
			row := c.row + step[i][0]
			col := c.col + step[i][1]
			if row >= 0 && row < m && col >= 0 && col < n && visited[row][col] != true {
				visited[row][col] = true
				ret += max(0, (c.height - heightMap[row][col]))
				heap.Push(&p, NewCell(row, col, max(c.height, heightMap[row][col])))
			}
		}
	}

	return ret
}
