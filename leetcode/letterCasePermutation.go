package main

import "fmt"

func main() {
	fmt.Println(letterCasePermutation("a1b1"))
}

func letterCasePermutation(S string) []string {
	return casePermutation(S, 0, "", []string{})
}

func casePermutation(S string, index int, str string, ret []string) []string {
	for index < len(S) && S[index] >= '1' && S[index] <= '9' {
		str += string(S[index])
		index++
	}
	if index == len(S) {
		ret = append(ret, str)
		return ret
	}
	var ret1, ret2 []string
	if S[index] >= 'a' && S[index] <= 'z' {
		ret1 = casePermutation(S, index+1, str + string(S[index] - 32), ret)
	} else {
		ret1 = casePermutation(S, index+1, str + string(S[index] + 32), ret)
	}
	ret2 = casePermutation(S, index+1, str + string(S[index]), ret)
	ret = append(ret, ret1...)
	ret = append(ret, ret2...)
	return ret
}
