package main

import "fmt"

func main() {
	fmt.Println(commonChars([]string{"bella","label","roller"}))
}

func commonChars(A []string) []string {
	m := [26]map[int]int{}
	for i := 0; i < 26; i++ {
		m[i] = make(map[int]int)
	}
	for i := 0; i < len(A); i++ {
		count := [26]int{}
		for j := 0; j < len(A[i]); j++ {
			idx := A[i][j] - 'a'
			count[idx]++
			m[idx][count[idx]]++
		}
	}
	ret := []string{}
	for k, v := range m {
		for _, v1 := range v {
			if v1 == len(A) {
				ret = append(ret, string('a' + k))
			}
		}
	}
	return ret
}
