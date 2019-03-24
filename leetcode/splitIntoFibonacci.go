package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(splitIntoFibonacci("112"))
	//fmt.Println(splitIntoFibonacci("11235813"))
	//fmt.Println(splitIntoFibonacci("112358130"))
	//fmt.Println(splitIntoFibonacci("1101111"))
	//fmt.Println(splitIntoFibonacci("0123"))
	v := splitIntoFibonacci("3611537383985343591834441270352104793375145479938855071433500231900737525076071514982402115895535257195564161509167334647108949738176284385285234123461518508746752631120827113919550237703163294909")
	fmt.Println(v)
	v1, v2 := v[0], v[1]
	count := 0
	count += len(strconv.Itoa(v[0]))
	count += len(strconv.Itoa(v[1]))
	for i := 2; i < len(v); i++ {
		count += len(strconv.Itoa(v[i]))
		fmt.Println(v[i], v1 + v2, v[i] - v1 - v2)
		v1, v2 = v2, v[i]
	}
	fmt.Println("count:", count)
}

func splitIntoFibonacci(S string) []int {
	fmt.Println(len(S))
	if len(S) < 3 {
		return []int{}
	}
	var v1, v2 int64
	for i := 0; i < len(S); i++ {
		if len(S[:i+1]) > 1 && S[0] == '0' {
			break
		}
		v1, _ = strconv.ParseInt(S[:i+1], 10, 64)
		if v1 >= 2 << 31 {
			break
		}
		for j := i + 1; j < len(S); j++ {
			if len(S[i+1:j+1]) > 1 && S[i+1] == '0' {
				break
			}
			v2, _ = strconv.ParseInt(S[i+1:j+1], 10, 64)
			if v2 >= 2 << 31 {
				break
			}
			if v1 + v2 >= 2 << 31 {
				break
			}
			ret, ok := sif(S[j+1:], v1, v2)
			if ok {
				return append([]int{int(v1), int(v2)}, ret...)
			}
		}
	}
	return []int{}
}

func sif(s string, v1, v2 int64) ([]int, bool) {
	if len(s) == 0 {
		return []int{}, true
	}
	var v3 int64
	for i := 0; i < len(s); i++ {
		if len(s[:i+1]) > 1 && s[0] == '0' {
			break
		}
		v3, _ = strconv.ParseInt(s[:i+1], 10, 64)
		if v3 >= 2 << 31 {
			break
		}
		if v3 == (v1 + v2) {
			ret, ok := sif(s[i+1:], v2, v3)
			if ok {
				//fmt.Println(s[i+1:], v2, v3)
				return append([]int{int(v3)}, ret...), true
			}
			return []int{}, false
		}
		if v3 > (v1 + v2) {
			break
		}
	}
	return []int{}, false
}
