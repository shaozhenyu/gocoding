package main

import (
	"fmt"
	"math"
)

func main() {
	p := [][]int{[]int{3,1},[]int{1,1},[]int{0,1},[]int{2,1},[]int{3,3},[]int{3,2},[]int{0,2},[]int{2,3}}
	p = [][]int{[]int{0,1},[]int{1,0},[]int{3,2},[]int{2,3},[]int{0,3},[]int{1,1},[]int{3,3},[]int{0,2}}
	fmt.Println(minAreaFreeRect(p))
}

func minAreaFreeRect(points [][]int) float64 {
	m := make(map[int][][]int)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			square := (points[j][0] - points[i][0]) * (points[j][0] - points[i][0]) + (points[j][1] - points[i][1]) * (points[j][1] - points[i][1])
			if _, ok := m[square]; !ok {
				m[square] = make([][]int, 0)
			}
			m[square] = append(m[square], []int{i, j})
		}
	}
	min := float64(2 << 31)
	for k, v := range m {
		if len(v) > 1 {
			for i := 0; i < len(v); i++ {
				for j := i + 1; j < len(v); j++ {
					if v[i][0] == v[j][0] || v[i][0] == v[j][1] || v[i][1] == v[j][0] || v[i][1] == v[j][1] {
						continue
					}
					ok, sq := checkRect(k, points[v[i][0]], points[v[i][1]], points[v[j][0]], points[v[j][1]])
					fmt.Println(v[i][0], v[i][1], v[j][0], v[j][1])
					if ok {
						fmt.Println(sq)
						fmt.Println(points[v[i][0]], points[v[i][1]], points[v[j][0]], points[v[j][1]])
					}
					if ok && sq < min {
						min = sq
					}
				}
			}
		}
	}
	fmt.Println(m)
	if min == float64(2 << 31) {
		min = 0.0
	}
	return min
}

func checkRect(t int, x1, x2, y1, y2 []int) (bool, float64) {
	fmt.Println(x1, x2, y1, y2)
	x1y1 := getSquare(y1[0] - x1[0], y1[1] - x1[1])
	x2y1 := getSquare(y1[0] - x2[0], y1[1] - x2[1])
	x1y2 := getSquare(y2[0] - x1[0], y2[1] - x1[1])
	x2y2 := getSquare(y2[0] - x2[0], y2[1] - x2[1])
	fmt.Println("ttttt : ", t, x1y1, x2y1, x1y2, x2y2)
	if (x1y1 + t == x2y1) && (x2y2 + t == x1y2) {
		return true, math.Sqrt(float64(t)) * math.Sqrt(float64(x1y1))
	}
	if (x1y2 + t == x2y2) && (x2y1 + t == x1y1) {
		return true, math.Sqrt(float64(t)) * math.Sqrt(float64(x1y2))
	}
	return false, 0.0
}

func getSquare(x, y int) int {
	return x * x + y * y
}
