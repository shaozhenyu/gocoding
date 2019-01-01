package main

import "fmt"

func main() {
	boxes := []int{1,3,2,2,2,3,3,1,2,2,3}
	boxes = []int{1,3,2,2,2,3,3,1,2,2,3,2,56,23,3,1,43,5,67,4,3,2,1,2,1,2,1,4,5,1,2,2,1,2,2,1,2,1}
	//boxes = []int{8,2,2,9,6,4,3,10,2,10,10,1,10,2,9,5,2,9,7,4,10,2,3,8,3,6,10,9,7,10,9,7,5,3,4,1,3,10,2,6,1,1,1,2,5,5,10,8,9,9}
	fmt.Println(removeBoxes(boxes))
}


func removeBoxes(boxes []int) int {
	if len(boxes) == 0 {
		return 0
	}
	ret := 0
	ret, boxes = delOneBox(boxes)
	m := make(map[int][]int)
	for i := 0; i < len(boxes); i++ {
		if _, ok := m[boxes[i]]; !ok {
			m[boxes[i]] = make([]int, 0, 100)
		}
		m[boxes[i]] = append(m[boxes[i]], i)
	}
	if len(m) == 1 {
		return ret + len(boxes) * len(boxes)
	}
	for _, v := range m {
		t := make([]int, len(boxes[:v[0]]))
		copy(t, boxes[:v[0]])
		t = append(t,  boxes[v[len(v)-1]+1:]...)
		ret1 := len(v) * len(v) + removeBoxes(t)
		start := v[0] + 1
		i := 1
		for ; i < len(v); i++ {
			ret1 += removeBoxes(boxes[start:v[i]])
			for i < len(v) - 1 && v[i] == v[i+1] {
				i++
			}
			start = v[i] + 1
		}
		ret = max(ret1, ret)
	}
	return ret
}

func delOneBox(boxes []int) (int, []int) {
	// fmt.Println("kkkkk", boxes)
	ret := 0
	m := make(map[int]int)
	for i := 0; i < len(boxes); i++ {
		m[boxes[i]]++
	}
	for i := 0; i < len(boxes); i++ {
		if m[boxes[i]] == 1 {
			boxes = append(boxes[:i], boxes[i+1:]...)
			i--
			ret++
		}
	}
	// fmt.Println("xxxxx", boxes)
	return ret, boxes
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
