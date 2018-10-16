package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "DoAB12(C(AoB)12)10"
	fmt.Println(countOfAtoms(s))
}

func countOfAtoms(formula string) string {
	m, idx := parse(formula, 0)
	fmt.Println(idx)
	fmt.Println(m)
	return split(m)
}

func split(m map[string]int) string {
	strs := make([]string, len(m))
	idx := 0
	for k, v := range m {
		val := ""
		if v > 1 {
			val = fmt.Sprintf("%s%d", k, v)
		} else {
			val = k
		}
		strs[idx] = val
		idx++
	}
	sort.Strings(strs)
	ret := ""
	for i := 0; i < len(strs); i++ {
		ret += strs[i]
	}
	return ret
}

func parse(formula string, idx int) (map[string]int, int) {
	m := make(map[string]int)
	if len(formula) == 0 {
		return m, idx
	}
	for idx < len(formula) {
		v := formula[idx]
		if v >= 'A' && v <= 'Z' {
			small := ""
			j := idx + 1
			for j < len(formula) && formula[j] >= 'a' && formula[j] <= 'z' {
				small += string(formula[j])
				j++
			}
			str := string(v) + small
			count := 0
			for j < len(formula) && formula[j] >= '0' && formula[j] <= '9' {
				count = count * 10 + int(formula[j] - '0')
				j++
			}
			if count == 0 {
				count = 1
			}
			m[str] += count
			idx = j
		} else if v == '(' {
			var m1 map[string]int
			m1,idx = parse(formula, idx+1)
			for k, v := range m1 {
				m[k] += v
			}
		} else {
			j := idx + 1
			count := 0
			for j < len(formula) && formula[j] >= '0' && formula[j] <= '9' {
				count = count * 10 + int(formula[j] - '0')
				j++
			}
			if count != 0 {
				for k, v := range m {
					m[k] = v * count
				}
			}
			idx = j
			return m, idx
		}
	}
	return m, idx
}
