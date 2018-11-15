package main

import "fmt"

func main() {
	s := "abpcplea"
	d := []string{"ale","apple","monkey","plea"}

	s = "aewfafwafjlwajflwajflwafj"
	d = []string{"apple","ewaf","awefawfwaf","awef","awefe","ewafeffewafewf"}
	fmt.Println(findLongestWord(s, d))
}

func findLongestWord(s string, d []string) string {
	m := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = make([]int, 0)
		}
		m[s[i]] = append(m[s[i]], i)
	}

	max := 0
	ret := ""
	for i := 0; i < len(d); i++ {
		str := d[i]
		if len(str) < max || (len(str) == max && str > ret) {
			continue
		}
		idx := -1
		j := 0
		for ; j < len(str); j++ {
			if v, ok := m[str[j]]; ok {
				x := 0
				for ; x < len(v); x++ {
					if v[x] > idx {
						idx = v[x]
						break
					}
				}
				if x == len(v) {
					break
				}
			} else {
				break
			}
		}
		if j == len(str) {
			if j > max {
				ret = str
				max = j
			} else if j == max {
				if ret > str {
					ret = str
				}
			}
		}

	}
	return ret
}
