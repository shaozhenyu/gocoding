package main

import (
	"fmt"
	"math"
)

func main() {
	//a := "goddtitngs"
	//b := "gtd"
	//a := "godding"
	//b := "godding"
	a := "rtmdx"
	b := "dmrtx"
	fmt.Println(findRotateSteps(a, b))

}

func minOf(a [][]int) int {
	min := math.MaxInt32

	for _, n := range a {
		if n[1] < min {
			min = n[1]
		}
	}

	return min
}

func distance(length int, i int, j int) int {
	a := int(math.Abs(float64(i - j)))
	b := int(math.Abs(float64(length - a)))

	if a < b {
		return a
	}

	return b
}

func findRotateSteps(ring string, key string) int {
	hash := make(map[rune][]int)

	for i, r := range ring {
		_, ok := hash[r]

		if ok {
			hash[r] = append(hash[r], i)
		} else {
			hash[r] = []int{i}
		}

	}

	result := [][]int{[]int{0, 0}}

	for _, ch := range key {
		current := hash[ch]
		next := make([][]int, len(current))

		for i := 0; i < len(current); i++ {
			min := math.MaxInt32

			for j := 0; j < len(result); j++ {
				dist := distance(len(ring), result[j][0], current[i]) + result[j][1]
				if dist < min {
					min = dist
				}
			}

			next[i] = []int{current[i], min}
		}

		result = next
	}

	return minOf(result) + len(key)
}

// godding
func findRotateSteps1(ring string, key string) int {
	fmt.Println("ring:", ring, len(ring))
	fmt.Println("key:", key)

	ret := 0
	lenK := len(key)
	lenR := len(ring)
	m := make(map[string][]int)
	for i := 0; i < len(ring); i++ {
		str := string(ring[i])
		if _, ok := m[str]; ok {
			m[str] = append(m[str], i)
		} else {
			m[str] = []int{i}
		}
	}
	py := 0
	for i := 0; i < lenK; i++ {
		val := string(key[i])
		s := m[val]
		p := lenR
		pos := 0
		fmt.Println("s:", s)
		for _, s1 := range s {
			right := abs(s1+py) % lenR
			left := abs(lenR-s1-py) % lenR
			fmt.Println("right:", right, "left:", left)
			if p > min(right, left) {
				fmt.Println("11111")
				if right <= left {
					pos = -1 * right
					p = right
				} else {
					pos = left
					p = left
				}
			}
		}
		fmt.Println("str : ", val, "p:", p, "pos:", pos)
		py += pos
		ret += p
		fmt.Println("ret:", ret, "py:", py)
	}

	return len(key) + ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
