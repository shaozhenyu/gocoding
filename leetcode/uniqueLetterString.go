package main

import "fmt"

func main() {
	fmt.Println(uniqueLetterString("ABC"), 10)
	fmt.Println(uniqueLetterString("ABA"), 8)
	fmt.Println(uniqueLetterString("ABCBBC"), 26)
}

func uniqueLetterString(S string) int {
	if len(S) == 0 {
		return 0
	}
	count := 0
	mod := 1000000007
	existsCount := [26]int{}
	exists := make(map[byte]int)
	for i := len(S) - 1; i >= 0; i-- {
		nextIdx, ok := exists[S[i]]
		if !ok {
			existsCount[S[i] - 'A'] = len(S) - i
		} else {
			existsCount[S[i] - 'A'] = nextIdx - i
		}
		for j := 0; j <= 25; j++ {
			count = (count + existsCount[j])%mod
		}
		exists[S[i]] = i
	}
	return count
}

func uniqueLetterString1(S string) int {
	if len(S) == 0 {
		return 0
	}
	count := 1
	preCount := 1
	mod := 1000000007
	exists := make(map[byte]int)

	exists[S[len(S) - 1]] = len(S) - 1
	for i := len(S) - 2; i >= 0; i-- {
		nowCount := 0
		nextIdx, ok := exists[S[i]]
		if !ok {
			nowCount = (len(S) - i + preCount)%mod
		} else {
			nowCount = (preCount - (len(S) - nextIdx) + nextIdx - i)%mod
		}
		fmt.Println("xx:", string(S[i]), i, nowCount)
		count = (count + nowCount)%mod
		preCount = nowCount
		exists[S[i]] = i
	}
	return count
}
