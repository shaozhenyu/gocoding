package main

import (
	"fmt"
)

func main() {
	points := []Point{Point{0, 2}, Point{1, 1}, Point{2, 2}, Point{2, 4}, Point{4, 2}, Point{3, 3}}
	fmt.Println(outerTrees(points))
}

type Point struct {
	X int
	Y int
}

func outerTrees(points []Point) []Point {
	if len(points) <= 3 {
		return points
	}
	m := make(map[int]bool)
	tuBao := func(p1, p2, p3 Point) int {
		return p1.X*p2.Y + p2.X*p3.Y + p3.X*p1.Y - p3.X*p2.Y - p2.X*p1.Y - p1.X*p3.Y
	}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if m[i] && m[j] {
				continue
			}
			left, right := 0, 0
			line := make([]int, 0)
			for k := 0; k < len(points); k++ {
				if k == i || k == j {
					continue
				}
				t := tuBao(points[i], points[j], points[k])
				if t > 0 {
					left++
				} else if t < 0 {
					right++
				} else {
					line = append(line, k)
				}
				if left != 0 && right != 0 {
					break
				}
			}
			if left == 0 || right == 0 {
				m[i] = true
				m[j] = true
				for _, l := range line {
					m[l] = true
				}
			}
		}
	}
	r := make([]Point, 0, len(points))
	for k := range m {
		r = append(r, points[k])
	}
	return r
}
