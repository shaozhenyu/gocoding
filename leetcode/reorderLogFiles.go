package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	fmt.Println(reorderLogFiles([]string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"}))
}

func reorderLogFiles(logs []string) []string {
	as := make([]string, 0)
	ds := make([]string, 0)
	for _, log := range logs {
		idx := strings.Index(log, " ")
		if log[idx+1] >= '0' && log[idx+1] <= '9' {
			ds = append(ds, log)
		} else {
			as = append(as, log)
		}
	}
	sort.Slice(as, func(i, j int) bool {
		idxI := strings.Index(as[i], " ")
		idxJ := strings.Index(as[j], " ")
		return as[i][idxI+1:] < as[j][idxJ+1:]
	})
	return append(as, ds...)
}
