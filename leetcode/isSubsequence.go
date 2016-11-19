package main

func main() {
	if isSubsequence("aa", "aba") {
		println("aaa")
	}
}

func isSubsequence(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)
	if lenS > lenT {
		return false
	}

	if lenS == 0 && lenT == 0 {
		return true
	}

	j := 0
	for i := 0; i < lenS; i++ {
		for j < lenT {
			if s[i] == t[j] {
				j++
				break
			} else {
				j++
			}
		}
		if j >= lenT && s[i] != t[j-1] {
			return false
		}
	}
	return true
}
