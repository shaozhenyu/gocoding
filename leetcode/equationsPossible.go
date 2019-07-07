package main

import "fmt"

func main() {
	fmt.Println(equationsPossible([]string{"c==c","b==d","x!=z"}))
	fmt.Println(equationsPossible([]string{"a==b","e==c","b==c","a!=e"}))
}

func equationsPossible(equations []string) bool {
	equal := [26]int{}
	for i := 0; i < len(equations); i++ {
		e := equations[i]
		x, y := int(e[0] - 'a'), int(e[3] - 'a')
		if e[1] == '=' {
			if equal[x] == 0 && equal[y] == 0 {
				equal[x] = x + 1
				equal[y] = x + 1
			} else if equal[x] != 0 && equal[y] == 0 {
				equal[y] = equal[x]
			} else if equal[x] == 0 && equal[y] != 0 {
				equal[x] = equal[y]
			} else {
				ey := equal[y]
				for t := 0; t < 26; t++ {
					if equal[t] == ey {
						equal[t] = equal[x]
					}
				}
			}
			fmt.Println(equal, equations[i])
		}
	}
	fmt.Println(equal)
	for i := 0; i < len(equations); i++ {
		e := equations[i]
		x, y := int(e[0] - 'a'), int(e[3] - 'a')
		if e[1] == '!' {
			if x == y {
				return false
			}
			if equal[x] == 0 || equal[y] == 0 {
				continue
			}
			if equal[x] == equal[y] {
				return false
			}
		}
	}
	return true
}
