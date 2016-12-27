package main

func main() {
	println(countSegments("azAZlo'aa  sd as a"))
}

func countSegments(s string) int {
	ret := 0
	isAlphabe := 0
	for i := 0; i < len(s); i++ {
		if (string(s[i]) == "'" || string(s[i]) == "-") && isAlphabe == 1 {
			continue
		}
		if (s[i] >= 65 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122) {
			if isAlphabe == 0 {
				ret++
			}
			isAlphabe = 1
		} else {
			isAlphabe = 0
		}
	}
	return ret
}
