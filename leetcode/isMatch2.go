package main

func main() {
	println("aa c*a*b*aa except ok :")
	if isMatch("aa", "c*a*b*aa") {
		println(" ok")
	} else {
		println(" flase")
	}

	println("a ab*a except false :")
	if isMatch("a", "ab*a") {
		println(" ok")
	} else {
		println("false")
	}

	println("a .*..a* except false :")
	if isMatch("a", ".*..a*") {
		println(" ok")
	} else {
		println("false")
	}
}

//'.' Matches any single character.
//'*' Matches zero or more of the preceding element.
func isMatch(s string, p string) bool {
	lenS := len(s)
	lenP := len(p)
	i, j := 0, 0

	for i < len(s) && j < len(p) {
		if s[i] == p[j] || string(p[j]) == "." {
			i++
			j++
		} else {
			
		}
	}
}
