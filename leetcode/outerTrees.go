package main

import (
	"fmt"
	"sort"
)

func main() {
	points := []Point{Point{1, 1}, Point{2, 2}, Point{2, 0}, Point{2, 4}, Point{3, 3}, Point{4, 2}}
	fmt.Println(outerTrees(points))
}

type Point struct {
	X int
	Y int
}

func tuBao(p1, p2, p3 Point) int {
	return  (p1.X*p2.Y + p2.X*p3.Y + p3.X*p1.Y - p3.X*p2.Y - p2.X*p1.Y - p1.X*p3.Y)
}

func divideLeft(points []Point, left, right int, outerM map[int]bool) {
	max := 0
	maxPoint := -1
	onLine := make(map[int]bool)
	if left < right {
		for i := left + 1; i < right; i++ {
			v := tuBao(points[left], points[i], points[right])
			if v > max {
				max = v
				maxPoint = i
			}
			if v == 0 {
				onLine[i] = true
			}
		}
	} else {
		for i := left - 1; i > right; i-- {
			v := tuBao(points[left], points[i], points[right])
			if v > max {
				max = v
				maxPoint = i
			}
			if v == 0 {
				onLine[i] = true
			}
		}
	}

	if maxPoint != -1 {
		outerM[maxPoint] = true
		divideLeft(points, left, maxPoint, outerM)
		divideLeft(points, maxPoint, right, outerM)
	} else {
		for k := range onLine {
			outerM[k] = true
		}
	}
}

// 分治法 44ms
func outerTrees(points []Point) []Point {
	sort.Slice(points, func(i, j int) bool {
		if points[i].X == points[j].X {
			return points[i].Y < points[j].Y
		}
		return points[i].X < points[j].X
	})
	outerM := make(map[int]bool)
	left, right := 0, len(points) - 1
	outerM[left] = true
	outerM[right] = true
	divideLeft(points, left, right, outerM)
	divideLeft(points, right, left, outerM)
	r := make([]Point, 0, len(points))
	for k := range outerM {
		r = append(r, points[k])
	}
	return r
}

// 穷举法 超时
func outerTrees5(points []Point) []Point {
	m := make(map[int]bool)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if m[i] && m[j] {
				continue
			}
			left, right := 0, 0
			sk := make([]int, 0)
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
					sk = append(sk, k)
				}
				if left != 0 && right != 0 {
					break
				}
			}
			if left == 0 || right == 0 {
				m[i] = true
				m[j] = true
				for _, skk := range sk {
					m[skk] = true
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
