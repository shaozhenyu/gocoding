// msd 高位优先的字符串排序 不定长
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
		"1oo1",
		"1",
		"1o",
	}
	fmt.Println(msd(s))
}

func msd(strs []string) []string {
	if len(strs) == 0 {
		return strs
	}
	return sort(strs, 0, len(strs)-1, 0)
}

func charAt(str string, d int) int {
	if d < len(str) {
		return int(str[d])
	}
	return -1
}

func sort(strs []string, lo, hi, d int) []string {
	if lo >= hi {
		return strs
	}
	R := 256
	n := len(strs)
	count := make([]int, R + 2)
	for i := lo; i <= hi; i++ {
		count[charAt(strs[i], d) + 2]++
	}
	for i := 1; i < R + 2; i++ {
		count[i] += count[i-1]
	}
	tmp := make([]string, n)
	for i := lo; i <= hi; i++ {
		tmp[count[charAt(strs[i], d) + 1]] = strs[i]
		count[charAt(strs[i], d) + 1]++
	}
	for i := lo; i <= hi; i++ {
		strs[i] = tmp[i - lo]
	}
	for r := 0; r <= R; r++ {
		sort(strs, lo + count[r], lo + count[r+1] - 1, d + 1)
	}
	return strs
}
