package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := "105"
	target := 5
	num = "123456789"
	target = 45
	fmt.Println(len(addOperators(num, target)))
}

func addOperators(num string, target int) []string {
	return addOperator(num, target, 1, false, 0, false)
}

func addOperator(num string, target int,
	preMul int, preMulFlag bool,
	preDel int, preDelFlag bool) []string {
	ret := []string{}
	if len(num) == 0 {
		return ret
	}
	m := make(map[string]struct{})
	for i := 0; i < len(num); i++ {
		if num[0] == '0' && i > 0 {
			break
		}
		v, _ := strconv.Atoi(string(num[:i+1]))
		if preDelFlag {
			ret = aos(num, fmt.Sprintf("%d+", v), i+1, target-(preDel-v*preMul), ret, 1, false, 0, false)
			ret = aos(num, fmt.Sprintf("%d-", v), i+1, target, ret, 1, false, preDel-v*preMul, true)
		} else {
			ret = aos(num, fmt.Sprintf("%d+", v), i+1, target-v*preMul, ret, 1, false, 0, false)
			ret = aos(num, fmt.Sprintf("%d-", v), i+1, target, ret, 1, false, v*preMul, true)
		}
		ret = aos(num, fmt.Sprintf("%d*", v), i+1, target, ret, preMul*v, true, preDel, preDelFlag)
	}
	for i := 0; i < len(ret); i++ {
		m[ret[i]] = struct{}{}
	}
	ret = []string{}
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

func aos(num string, pre string, idx, target int, ret []string,
	preMul int, preMulFlag bool,
	preDel int, preDelFlag bool) []string {
	if idx == len(num) {
		if target == 0 && !preMulFlag && !preDelFlag {
			ret = append(ret, pre[:len(pre)-1])
		}
		return ret
	} else {
		vals := addOperator(num[idx:], target, preMul, preMulFlag, preDel, preDelFlag)
		if len(vals) > 0 {
			for _, val := range vals {
				ret = append(ret, pre+val)
			}
		}
	}
	return ret
}
