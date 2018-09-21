package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceWords([]string{}, "aa bb"))
}

func replaceWords(dict []string, sentence string) string {
	m := make(map[string]bool)
	for _, d := range dict {
		m[d] = true
	}
	ret := ""
	sli := strings.Split(sentence, " ")
	Loop:
	for _, s := range sli {
		for i := 0; i < len(s); i++ {
			if v, ok := m[s[:i+1]]; ok {
				ret += fmt.Sprintf("%s ", v)
				continue Loop
			}
		}
		ret += fmt.Sprintf("%s ", s)
	}
	return ret[:len(ret)-1]
}
