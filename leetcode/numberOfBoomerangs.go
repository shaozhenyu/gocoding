package main

func main() {
	a := [][]int{{0,0}, {0,1}, {0,2}}
	println(numberOfBoomerangs(a))
}

func numberOfBoomerangs(point [][]int) int {
	size := len(point)
	if size <= 1 {
		return 0
	}
	ret := 0
	for i := 0; i < size; i++ {
		pp := point[i]
		m := make(map[int]int, size)
		for j := 0; j < size; j++ {
			distance := calDistance(pp, point[j])
			if _, ok := m[distance]; ok {
				m[distance] += 1
			} else {
				m[distance] = 1
			}
		}
		for _, v := range m {
			ret += v * (v-1)
		}
	}
	return ret
}

func calDistance(p1, p2 []int) int {
	a := (p1[0] - p2[0]) * (p1[0] - p2[0])
	b := (p1[1] - p2[1]) * (p1[1] - p2[1])
	return a + b
}
