package main

import "fmt"

func main() {

}

type STree struct {
	left, right int
	length int
	flag int
}

type Node struct {
	x1, x2 int
	y int
	flag int
}

func build(st []STree, idx, left, right int) {
	st[idx] = STree{left, right, 0, 0}
	if left == right {
		return
	}
	mid := left + (right - left)/2
	build(st, idx << 1 + 1, left, mid)
	build(st, idx << 1 + 2, mid+1, right)
}

func pushUp(st []STree, x []int, idx int) {
	if st[idx].flag != 0 {
		st[idx].length = x[st[idx].right + 1] - x[st[idx].left]
	} else if st[idx].left == st[idx].right {
		st[idx].length = 0
	} else {
		st[idx].length = st[idx << 1 + 1].length + st[idx << 1 + 2].length
	}
}

func updateSTree(st []STree, x []int, idx, left, right, flag int) {
	if left > st[idx].right || right < st[idx].left {
		return
	}
	if left <= st[idx].left && right >= st[idx].right {
		st[idx].flag += flag
		pushUp(st, x, idx)
		return
	}
	updateSTree(st, x, idx << 1 + 1, left, right, flag)
	updateSTree(st, x, idx << 1 + 2, left, right, flag)
	pushUp(st, x, idx)
}

// rectangle[i] = [x1, y1, x2, y2]
func rectangleArea(rectangles [][]int) int {
	size := len(rectangles) * 2
	x := make([]int, size)
	node := make([]Node, size)
	idx := 0
	for i := 0; i < len(rectangles); i++ {
		x1, y1, x2, y2 := rectangles[i][0], rectangles[i][1], rectangles[i][2], rectangles[i][3]
		x[idx], x[idx+1] = x1, x2
		node[idx] = Node{x1, x2, y1, 1}
		node[idx+1] = Node{x1, x2, y2, -1}
		idx += 2
	}
	sort.Ints(x)
	sort.Slice(node, func(i, j int) bool {
		return node[i].y < node[j].y
	})

	pos := 1
	x[0] = x[0]
	for i := 1; i < size; i++ {
		if x[i] != x[i-1] {
			x[pos] = x[i]
			pos++
		}
	}
	stree := make([]STree, 2 * int(math.Pow(2, math.Ceil(math.Log2(float64(size))))) - 1)
	build(stree, 0, 0, size-1)
	var ret int64
	for i := 0; i < size - 1; i++ {
		left := getTreeIdx(x, 0, pos-1, node[i].x1)
		right := getTreeIdx(x, 0, pos-1, node[i].x2) - 1
		updateSTree(stree, x, 0, left, right, node[i].flag)
		ret += int64(stree[0].length * (node[i+1].y - node[i].y))
	}
	return int(ret%1000000007)
}

func getTreeIdx(x []int, left, right, val int) int {
	var mid = 0
	for left <= right {
		mid = left + (right - left)/2
		if x[mid] < val {
			left = mid + 1
		} else if x[mid] > val {
			right = mid - 1
		} else {
			break
		}
	}
	return mid
}
