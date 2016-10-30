package main

import "fmt"
import "strconv"

func main() {
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	var m, n string
	if len(a) >= len(b) {
		m = a
		n = b
	} else {
		m = b
		n = a
	}

	i := len(m) - 1
	j := len(n) - 1
	add := 0
	str := ""
	for ; i >= 0; i-- {
		if j < 0 {
			if add == 1 {
				if m[i] == 49 {
					str = "0" + str
					add = 1
				} else {
					str = "1" + str
				}
			} else {
				str = string(m[i]) + str
			}
		} else {
			sum := m[i] + n[j]
			if sum == 98 {
				if add == 1 {
					str = "1" + str
				} else {
					str = "0" + str
				}
				add = 1
			} else if sum == 97 {
				if add == 1 {
					str = "0" + str
					add = 1
				} else {
					str = "1" + str
					add = 0
				}
			} else if sum == 96 {
				str = strconv.Itoa(add) + str
				add = 0
			} else {
				return ""
			}
			j--
		}
	}

	if add == 1 {
		str = "1" + str
	}

	return str
}
