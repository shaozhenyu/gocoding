package main

import (
	"fmt"
)

func main() {
	var a, b string
	var t int
	a = "dinitrophenylhydrazine"
	b = "acetylphenylhydrazine"
	t = minDistance(a, b)
	if t == 11 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	a = "intention"
	b = "execution"
	t = minDistance(a, b)
	if t == 8 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	a = "dinitrophenylhydrazine"
	b = "dimethylhydrazine"
	t = minDistance(a, b)
	if t == 9 {
		fmt.Println("true")
	} else {
		fmt.Println(t, 9)
		fmt.Println("false")
	}
	a = "pneumonoultramicroscopicsilicovolcanoconiosis"
	b = "ultramicroscopically"
	t = minDistance(a, b)
	if t == 29 {
		fmt.Println("true")
	} else {
		fmt.Println(t, 29)
		fmt.Println("false")
	}

	a = "dyuvukpqfrzklbafvcserjrfomjnujercedkqmyehwwcmgkbgpxouuatfjhnmnirowwutgnodtnhuqskcrudfzeuroaanjb"
	b = "uvukklcsjrfmjjkqmecpxouatfhnmwuthkcrufeoj"
	t = minDistance(a, b)
	if t == 54 {
		fmt.Println("true")
	} else {
		fmt.Println(t, 54)
		fmt.Println("false")
	}
}

func minDistance1(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	dp := make([][]int, len1+1)

	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
		for j := 0; j <= len2; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				}
			}
		}
	}
	return len1 + len2 - 2*dp[len1][len2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {

	//先把前面相同的去掉
	for len(word1) != 0 && len(word2) != 0 && word1[0] == word2[0] {
		word1 = word1[1:]
		word2 = word2[1:]
	}

	w2 := make(map[byte][]int)
	len1, len2 := len(word1), len(word2)
	if len1 > len2 {
		word1, word2 = word2, word1
		len1, len2 = len2, len1
	}

	//把word2存入map中，key表示这个字符，value是一个数组，存放的是这个数组在该字符串中出现的位置
	for i := 0; i < len(word2); i++ {
		key := word2[i]
		if _, ok := w2[key]; ok {
			w2[key] = append(w2[key], i)
		} else {
			w2[key] = []int{i}
		}
	}

	//递归调用 example：word1:abcd， 分别比较abcd，bcd，cd，d和word2的最大相同数，取最大值。
	//在解abcd的时候 再调用函数对abcd的子串取最大值。大量的重复计算和递归
	max := getMaxSame(word1, 0, w2, -1)

	return len1 + len2 - 2*max
}

func getMaxSame(str string, n int, m map[byte][]int, prev int) int {
	max := 0
	lenS := len(str)
	for i := n; i < lenS; i++ {
		same := 0
		if v, ok := m[str[i]]; ok {
			for j := 0; j < len(v); j++ {
				//这个值的位置一定要大于上一个相同值
				if v[j] > prev {
					same = 1
					same += getMaxSame(str, i+1, m, v[j])
					break
				}
			}
		}
		if same > max {
			max = same
		}
		if max > lenS-i {
			break
		}
	}
	return max
}
