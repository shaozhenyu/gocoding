package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(compareVersion("0.1", "1.1"), -1)
	fmt.Println(compareVersion("1.0.1", "1"), 1)
	fmt.Println(compareVersion("7.5.2.4", "7.5.3"), -1)
	fmt.Println(compareVersion("1.01", "1.001"), 0)
	fmt.Println(compareVersion("1.0", "1.0.0"), 0)
}

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	i := 0
	for ; i < len(v1) && i < len(v2); i++ {
		vInt1, _ := strconv.Atoi(v1[i])
		vInt2, _ := strconv.Atoi(v2[i])
		if vInt1 > vInt2 {
			return 1
		} else if vInt1 < vInt2 {
			return -1
		}
	}
	for i < len(v1) {
		vInt1, _ := strconv.Atoi(v1[i])
		if vInt1 > 0 {
			return 1
		}
		i++
	}
	for i < len(v2) {
		vInt2, _ := strconv.Atoi(v2[i])
		if vInt2 > 0 {
			return -1
		}
		i++
	}
	return 0
}
