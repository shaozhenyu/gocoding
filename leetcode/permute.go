package main

import "fmt"


func main() {
	fmt.Println(permute("{a,b}c{d,e}f"))
}

var ret []string

func permute(S string) []string {
	ret = make([]string, 0)
	b := make([]byte, 0)
	more := make([][]byte, 0)
	for i := 0; i < len(S); {
		if S[i] == '{' {
			i++
			t := make([]byte, 0)
			for ; S[i] != '}'; i++ {
				if S[i] == ',' {
					continue
				}
				t = append(t, S[i])
			}
			more = append(more, t)
			b = append(b, '.')
		} else {
			b = append(b, S[i])
		}
		i++
	}
	if len(more) == 0 {
		return []string{S}
	}
	add := make([]int, 0)
	for i := 0; i < len(b); i++ {
		if b[i] == '.' {
			add = append(add, i)
		}
	}
	addOne(b, 0, add, more)
	return ret
}

func addOne(b []byte, idx int, add []int, more [][]byte) {
	if idx >= len(add) {
		ret = append(ret, string(b))
		return
	}
	for i := 0; i < len(more[idx]); i++ {
		b[add[idx]] = more[idx][i]
		addOne(b, idx+1, add, more)
		b[add[idx]] = '.'
	}
}
