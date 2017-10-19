package main

import (
	"fmt"
	"strings"
)

func main() {
	equations := [][]string{[]string{"a", "b"}, []string{"b", "c"}}
	values := []float64{2.0, 3.0}
	queriers := [][]string{[]string{"a", "c"}, []string{"b", "c"}, []string{"a", "e"}, []string{"a", "a"}, []string{"b", "c"}}
	fmt.Println(calcEquation(equations, values, queriers))
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	fmt.Println(equations, values, queries)
	m := make(map[string]float64)
	n := make(map[string]bool)
	var str, str1, str2 string
	for i := 0; i < len(values); i++ {
		a := equations[i][0]
		b := equations[i][1]
		n[a] = true
		n[b] = true
		str1 = a + "+" + b
		str2 = b + "+" + a
		m[str1] = values[i]
		m[str2] = 1.0 / values[i]
	}

	for i := 0; i < len(values); i++ {
		a := equations[i][0]
		b := equations[i][1]

		for key, value := range m {
			s := strings.Split(key, "+")
			c := s[0]
			d := s[1]
			if (a == c && b == d) || (a == d && b == c) {
				continue
			}
			if a == c {
				str1 = d + "+" + b
				str2 = b + "+" + d
				m[str1] = values[i] / value
				m[str2] = value / values[i]
			}

			if a == d {
				str1 = c + "+" + b
				str2 = b + "+" + c
				m[str1] = values[i] * value
				m[str2] = 1 / m[str1]
			}

			if b == c {
				str1 = a + "+" + d
				str2 = d + "+" + a
				m[str1] = values[i] * value
				m[str2] = 1 / m[str1]
			}

			if b == d {
				str1 = a + "+" + c
				str2 = c + "+" + a
				m[str1] = values[i] / value
				m[str2] = value / values[i]
			}
		}
	}

	fmt.Println(m)

	ret := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		str = queries[i][0] + "+" + queries[i][1]
		if v, ok := m[str]; ok {
			ret[i] = v
		} else {
			if queries[i][0] == queries[i][1] {
				if _, ok := n[queries[i][0]]; ok {
					ret[i] = 1.0
				}
			} else {
				ret[i] = -1.0
			}
		}
	}

	return ret
}
