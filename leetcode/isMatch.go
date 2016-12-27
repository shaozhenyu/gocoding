package main

func main() {
	println("aa c*a*b*aa except ok :")
	if isMatch("aa", "c*a*b*aa") {
		println(" ok")
	}

	println("a ab*a except false :")
	if isMatch("a", "ab*a") {
		println(" ok")
	}

	println("a .*..a* except false :")
	if isMatch("a", ".*..a*") {
		println(" ok")
	}

	println("aaa ab*ac.*a except false :")
	if isMatch("aaa", "ab*ac.*a") {
		println(" ok")
	}

	println("abc ab*c*a*bc except true :")
	if isMatch("abc", "ab*c*a*bc") {
		println(" ok")
	}
}

//'.' Matches any single character.
//'*' Matches zero or more of the preceding element.
func isMatch(s string, p string) bool {
	i := len(s) - 1
	j := len(p) - 1

	for i >= 0 && j >= 0 {
		if (s[i] == p[j]) || (string(p[j]) == ".") {
			i--
			j--
		} else if string(p[j]) == "*" {
			if j > 0 && string(p[j-1]) == "." {
				j = j - 2
				continue
			}
			for (string(p[j]) == "*") && j >= 0 {
				j--
			}
			if j < 0 {
				return false
			}

			sameA := 0
			if (p[j] == s[i]) || (string(p[j]) == ".") {
				for i > 0 && s[i] == s[i-1] {
					sameA++
					i--
				}
				if i == 0 {
					break
				}
				if j == 0 {
					return false
				}
				i--
				j--
				tmp := j
				println("sameA: ", sameA)
				for j >= 0 && sameA > 0 {
					if string(p[j]) == "." || (p[j] == s[i+1]) {
						j--
						sameA--
						tmp = j
					} else if string(p[j]) != "*" {
						j = tmp
						break
					} else {
						tmp = j
						for string(p[j]) == "*" {
							j--
						}
					}
				}
			} else {
				j--
			}
		} else {
			return false
		}
	}

	if j < 0 && i >= 0 {
		return false
	}

	flag := 0
	for j >= 0 {
		if string(p[j]) != "*" && flag == 0 {
			println("aaaaa, ", j)
			t := j
			i, j = 0, 0
			for i <= t {
				if (s[i] == p[j]) {
					i++
					j++
				} else {
					return false
				}
			}
			return true
		} else if string(p[j]) != "*" && flag == 1 {
			j--
			flag = 0
		}
		for j >= 0 && string(p[j]) == "*" {
			flag = 1
			j--
		}
	}

	return true
}
