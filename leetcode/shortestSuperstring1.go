package main

import "fmt"

func main() {
	A := []string{"catg","ctaagt","gcta","ttca","atgcatc"}
	fmt.Println(shortestSuperstring(A), "gctaagttcatgcatc")
}

func shortestSuperstring(A []string) string {
	
	return ""
}

func connect(ft, bh string) string {
	for i := len(bh) - 1; i >= 0; i-- {
		if bh[i] == ft[len(ft)-1] {
			if len(ft) >= len(bh[:i+1]) && bh[:i+1] == ft[len(ft) - len(bh[:i+1]):] {
				return ft + bh[i+1:]
			}
		}
	}
	return ft + bh
}
