package main

import (
	"fmt"
	"sort"
)

func main() {
	// A := []string{"catg","ctaagt","gcta","ttca","atgcatc"}
	// fmt.Println(shortestSuperstring(A), "gctaagttcatgcatc")
	// fmt.Println(shortestSuperstring([]string{"orugbsuuxowmhjh","zjyxzmpduthlsioor","qtxocgehmhfqnstl","tlrlcnnrsyryfrywuebq","hozjyxzmpduthlsio","hjhdmnqtxocgehm","mjhzwdudlnbfkjawqacf","hfqnstlrlcnnrsyryfry","yfrywuebqhvwewzmq","zzieemjhzwdudlnbfkj","nnrsyryfrywuebqhvw","acfgaihbhozjyxzmpdut"}), "gctaagttcatgcatc")
	// fmt.Println(shortestSuperstring([]string{"abcde","cdef"}), "gctaagttcatgcatc")
	fmt.Println(shortestSuperstring([]string{"xas","nxv","ownx","xvf","vfv"}), "ownxvfvxas")
}

type Node struct {
	idx int
	size int
	add string
}

func shortestSuperstring(A []string) string {
	used := make([]bool, len(A))
	cm := make(map[[2]int]string)
	other := ""
	for i := 0; i < len(A); i++ {
		used[i] = true
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if i == j {
				continue
			}
			add := connect(A[i], A[j])
			cm[[2]int{i, j}] = add
			if len(add) < len(A[j]) {
				used[i], used[j] = false, false
			}
		}
	}
	newA := make([]string, 0, len(A))
	for i := 0; i < len(A); i++ {
		if used[i] {
			other += A[i]
		} else {
			newA = append(newA, A[i])
		}
	}
	return shortestS(newA, cm) + other
}

func shortestS(A []string, cm map[[2]int]string) string {
	short := ""
	shortLen := 10000
	connectM := make(map[int][]Node)
	used := make([]bool, len(A))
	for i := 0; i < len(A); i++ {
		connectM[i] = make([]Node, 0, len(A))
		for j := 0; j < len(A); j++ {
			add := cm[[2]int{i, j}]
			connectM[i] = append(connectM[i], Node{j, len(add), add})
		}
	}
	for _, v := range connectM {
		sort.Slice(v, func(i, j int) bool {
			return v[i].size < v[j].size
		})
	}
	for i := 0; i < len(A); i++ {
		used[i] = true
		backtrack(A, used, i, 1, A[i], &short, &shortLen, connectM)
		used[i] = false
	}
	return short
}

func backtrack(A []string, used []bool, idx, count int, now string, short *string, shortLen *int, conn map[int][]Node) {
	if count == len(A) {
		if len(now) < *shortLen {
			*short = now
			*shortLen = len(now)
		}
		return
	}
	if len(now) >= *shortLen {
		return
	}
	for _, n := range conn[idx] {
		i := n.idx
		add := n.add
		if used[i] {
			continue
		}
		used[i] = true
		backtrack(A, used, i, count + 1, now + add, short, shortLen, conn)
		used[i] = false
	}
}

func connect(ft, bh string) string {
	for i := len(bh) - 1; i >= 0; i-- {
		if bh[i] == ft[len(ft)-1] {
			if len(ft) >= len(bh[:i+1]) && bh[:i+1] == ft[len(ft) - len(bh[:i+1]):] {
				return bh[i+1:]
			}
		}
	}
	return bh
}
