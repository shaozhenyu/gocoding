package main

import "fmt"

func main() {
	fmt.Println(camelMatch([]string{"FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"}, "FB"))
}

func camelMatch(queries []string, pattern string) []bool {
	ret := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		p := pattern
		idx := 0
		j := 0
		for ; j < len(queries[i]); j++ {
			if idx == len(p) {
				if queries[i][j] >= 'A' && queries[i][j] <= 'Z' {
					break
				}
				continue
			}
			if queries[i][j] == p[idx] {
				idx++
			} else {
				if queries[i][j] >= 'A' && queries[i][j] <= 'Z' {
					break
				}
			}
		}
		if j == len(queries[i]) && idx == len(p) {
			ret[i] = true
		}
	}
	return ret
}
