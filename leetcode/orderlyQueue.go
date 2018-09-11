package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(orderlyQueue("casdasfafgdffdadsacbbddaa", 10))
}

func orderlyQueue(S string, K int) string {
	if K <= 0 {
		return S
	}
	sli := make([]string, len(S))
	ret := ""
	if K > 1 {
		for i := 0; i < len(S); i++ {
			sli[i] = string(S[i])
		}
		sort.Strings(sli)
		for i := 0; i < len(S); i++ {
			ret += sli[i]
		}
	} else {
		ret = S
		for i := 0; i < len(S); i++ {
			S = S[1:] + string(S[0])
			if S < ret {
				ret = S
			}
		}

	}
	return ret
}
