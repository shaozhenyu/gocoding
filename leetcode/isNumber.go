package main

import (
	"fmt"
	"strings"
)

func main() {
	if isNumber(" -. ") {
		fmt.Println("ok")
	}
	if isNumber(" +.8 ") {
		fmt.Println("ok")
	}
	if isNumber(" 13e.1 ") {
		fmt.Println("ok")
	}
	if isNumber(" 1.e+1 ") {
		fmt.Println("ok")
	}
	if isNumber(" 1.e-1 ") {
		fmt.Println("ok")
	}
}

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	println(s)

	length := len(s)
	if length == 0 {
		return false
	}

	flagO := 0
	flagP := 0
	flagE := 0

	if s[0] < 48 || s[0] > 57 {
		if length <= 1 {
			return false
		}
		if s[0] != 43 && s[0] != 45 { //- +
			if s[0] == 46 { // .
				flagP = 1
			} else {
				return false
			}
		} else {
			flagO = 1
		}
	}

	for i := 1; i < length; i++ {
		if i == 1 && (flagO == 1 || flagP == 1) {
			if s[i] < 48 || s[i] > 57 {
				if i + 1 >= length {
					return false
				}
				if s[i] == 46 && flagP == 0 {
					flagP = 2
				} else {
					return false
				}
			}
		} else if s[i] < 48 || s[i] > 57 {
			if s[i] == 46 {
				if flagP >= 1 || flagE == 1 {
					return false
				} else {
					flagP = 1
				}
			} else if s[i] == 101 {
				if flagE == 1 || (i+1 >= length) {
					return false
				} else if (s[i+1] == 43 || s[i+1] == 45) {
					i++
					if (i+1 >= length) {
						return false
					}
				} else {
					flagE = 1
				}
			} else {
				return false
			}
		}
	}
	return true
}
