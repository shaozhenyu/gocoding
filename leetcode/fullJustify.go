package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	words = []string{"What","must","be","acknowledgment","shall","be"}
	maxWidth := 16
	fmt.Println(fullJustify(words, maxWidth))
}

func fullJustify(words []string, maxWidth int) []string {
	ret := make([][]string, 0)
	lineCount := make([]int, 0)
	count := 0
	line := 0
	now := make([]string, 0)
	for i := 0; i < len(words); i++ {
		s := words[i]
		if count + len(s) <= maxWidth {
			now = append(now, s)
			count += len(s) + 1
			line += len(s)
		} else {
			ret = append(ret, now)
			lineCount = append(lineCount, line)
			count = len(s) + 1
			line = len(s)
			now = []string{s}
		}
	}

	res := make([]string, len(ret))
	for i := 0; i < len(ret); i++ {
		place := len(ret[i]) - 1
		space := maxWidth - lineCount[i]
		s := ""
		if place == 0 {
			s += ret[i][0] + strings.Repeat(" ", space)
			res[i] = s
			continue
		}
		a := space/place
		b := space%place
		j := 0
		for ; j < len(ret[i]) - 1; j++ {
			add := 0
			if b > 0 {
				add = 1
				b--
			}
			s += ret[i][j] + strings.Repeat(" ", a + add)
		}
		s += ret[i][j]
		res[i] = s
	}
	s := ""
	for i := 0; i < len(now); i++ {
		s += now[i] + " "
	}
	if len(s) > maxWidth {
		s = s[:maxWidth]
	} else {
		s += strings.Repeat(" ", maxWidth - len(s))
	}
	res = append(res, s)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
	return res
}
