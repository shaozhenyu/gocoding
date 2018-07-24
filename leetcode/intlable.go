package main

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

type IntLable struct {
	start int
	end int
}

func partitionLabels(S string) []int {
	if len(S) == 0 {
		return []int{}
	}
	b := []byte(S)
	m := make(map[byte]*IntLable)
	for i := 0; i < len(b); i++ {
		if _, ok := m[b[i]]; !ok {
			m[b[i]] = &IntLable{i, i}
		}
		m[b[i]].end = i
	}
	ret := []int{}
	start, end := 1, 0
	for i := 0; i < len(b); i++ {
		if start > end {
			start, end = m[b[i]].start, m[b[i]].end
		}
		if i == end && m[b[i]].end <= end {
			ret = append(ret, end - start + 1)
			start = i + 1
			continue
		}
		if i < end {
			if m[b[i]].end <= end {
				continue
			} else {
				end = m[b[i]].end
			}
		}
	}
	if start < end {
		ret = append(ret, end - start + 1)
	}
	return ret
}
