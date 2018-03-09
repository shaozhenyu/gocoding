package main

import "fmt"

func main() {
	fmt.Println(partitionLabels("aabbcccd"))
}

func partitionLabels(S string) []int {
	m := make(map[byte][]int)
	for i := 0; i < len(S); i++ {
		if _, ok := m[S[i]]; ok {
			m[S[i]][1] = i
		} else {
			m[S[i]] = []int{i, 0}
		}
	}

	ret := []int{}

	start, end := 0, 0
	for start < len(S) {
		end = m[S[start]][1]
		fmt.Println(start, end)
		if end == 0 {
			ret = append(ret, 1)
			start += 1
			continue
		}

		for i := start + 1; i < end; i++ {
			if m[S[i]][1] < end {
				continue
			} else {
				end = m[S[i]][1]
			}
		}
		ret = append(ret, end - start)
		if end >= len(S) - 1 {
			break
		}
		start = end + 1
	}
	return ret
}
