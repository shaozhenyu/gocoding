package main

import "fmt"

func main() {
	points := [][]int{[]int{1,1},[]int{1,3},[]int{3,1},[]int{3,3},[]int{2,2}}
	fmt.Println(minAreaRect(points))
}

func minAreaRect(points [][]int) int {
	mx := make(map[int][]int)
	my := make(map[int][]int)
	fmt.Println(points)
	for i := 0; i < len(points); i++ {
		x, y := points[i][0], points[i][1]
		if _, ok := mx[x]; !ok {
			mx[x] = make([]int, 0)
		}
		mx[x] = append(mx[x], i)
		if _, ok := my[y]; !ok {
			my[y] = make([]int, 0)
		}
		my[y] = append(my[y], i)
	}
	fmt.Println(mx)
	fmt.Println(my)
	min := 2 << 31
	for x, idxs := range mx {
		if len(idxs) < 2 {
			continue
		}
		for i := 0; i < len(idxs); i++ {
			iy := points[idxs[i]][1]
			if _, ok := my[iy]; !ok {
				continue
			}
			for j := i + 1; j < len(idxs); j++ {
				jy := points[idxs[j]][1]
				if _, ok := my[jy]; !ok {
					continue
				}
				for _, idx3 := range my[iy] {
					if idx3 == idxs[i] || idx3 == idxs[j] {
						continue
					}
					for _, idx4 := range my[jy] {
						if idx4 == idxs[i] || idx4 == idxs[j] {
							continue
						}
						if points[idx3][0] != points[idx4][0]{
							continue
						}
						fmt.Println(idxs[i], idxs[j], idx3, idx4)
						v := abs((points[idx3][0] - x) * (jy - iy))
						if v < min {
							min = v
						}
					}
				}
			}
		}
	}
	if min == 2 << 31 {
		min = 0
	}
	return min
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
