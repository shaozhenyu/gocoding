// LSD 低位优先的字符串排序 定长
package main

import "fmt"

func main() {
	s := []string{
		"bcd",
		"acd",
		"1oo",
		"abc",
		"1oo",
		"211",
	}
	fmt.Println(lsd(s))
}

func lsd(strs []string) []string {
	if len(strs) == 0 {
		return strs
	}
	R := 256
	w := len(strs[0])
	n := len(strs)
	for d := w - 1; d >= 0; d-- {
		count := make([]int, R)
		for i := 0; i < n; i++ {
			count[int(strs[i][d])+1]++
		}
		for i := 1; i < R; i++ {
			count[i] += count[i-1]
		}
		tmp := make([]string, n)
		for i := 0; i < n; i++ {
			tmp[count[strs[i][d]]] = strs[i]
			count[strs[i][d]]++
		}
		strs = tmp
	}
	return strs
}
