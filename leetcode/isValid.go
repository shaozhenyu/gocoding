package main

import "fmt"

func main() {
	if isValid("[()([()()])][({})]") {
		fmt.Println("111")
	} else {
		fmt.Println("222")
	}
}

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	arr := [100000]string{}

	arr[0] = string(s[0])
	tmp := 0

	for i := 1; i < len(s); i++ {
		a := arr[tmp]
		b := string(s[i])
		if ( a == "(" && b == ")") || (a == "[" && b == "]") || (a == "{" && b == "}") {
			if tmp == 0 {
				arr[tmp] = ""
				if (i + 1) >= len(s) {
					return true
				} else {
					arr[tmp] = string(s[i+1])
					i++
				}
			} else {
				arr[tmp] = ""
				tmp--
			}
		} else if (b == ")") || (b == "}") || (b == "]") {
			return false
		} else {
			tmp++
			arr[tmp] = b
		}
	}

	if string(arr[0]) != "" {
		return false
	}

	return true
}
