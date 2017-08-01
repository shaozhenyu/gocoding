package main

import (
	"fmt"
)

func main() {
	var a, b string
	var t int
	// a = "dinitrophenylhydrazine"
	// b = "acetylphenylhydrazine"
	// t = minDistance(a, b)
	// if t == 11 {
	// 	fmt.Println("true")
	// } else {
	// 	fmt.Println("false")
	// }
	// a = "intention"
	// b = "execution"
	// t = minDistance(a, b)
	// if t == 8 {
	// 	fmt.Println("true")
	// } else {
	// 	fmt.Println("false")
	// }

	// a = "dinitrophenylhydrazine"
	// b = "dimethylhydrazine"
	// t = minDistance(a, b)
	// if t == 9 {
	// 	fmt.Println("true")
	// } else {
	// 	fmt.Println(t, 9)
	// 	fmt.Println("false")
	a = "pneumonoultramicroscopicsilicovolcanoconiosis"
	b = "ultramicroscopically"
	t = minDistance(a, b)
	if t == 29 {
		fmt.Println("true")
	} else {
		fmt.Println(t, 29)
		fmt.Println("false")
	}
}

func minDistance(word1 string, word2 string) int {

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

	for i := 0; i < len(word2); i++ {
		key := word2[i]
		if _, ok := w2[key]; ok {
			w2[key] = append(w2[key], i)
		} else {
			w2[key] = []int{i}
		}
	}

	fmt.Println(w2)

	fmt.Println(word1)
	fmt.Println(word2)
	maxSame := 0
	for i := 0; i < len1; i++ {
		same := 0
		if v, ok := w2[word1[i]]; ok {
			//fmt.Println("start: ", string(word1[i]))
			same = 1
			same += getMaxSame(word1, i+1, w2, v[0])
		}
		//fmt.Println("end: ", same)
		if same > maxSame {
			maxSame = same
		}
	}
	fmt.Println(len1, len2, maxSame)

	return len1 + len2 - 2*maxSame
}

func getMaxSame(str string, n int, m map[byte][]int, prev int) int {
	max := 0
	for i := n; i < len(str); i++ {
		same := 0
		if v, ok := m[str[i]]; ok {
			for j := 0; j < len(v); j++ {
				if v[j] > prev {
					same = 1
					same += getMaxSame(str, i, m, v[j])
					break
				}
			}
		}
		if same > max {
			max = same
		}
	}
	return max
}
