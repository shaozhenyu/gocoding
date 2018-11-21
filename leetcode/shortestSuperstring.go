package main

import "fmt"

// Output: "gctaagttcatgcatc"
func main() {
	A := []string{"catg","ctaagt","gcta","ttca","atgcatc"}
	fmt.Println(A)
	fmt.Println(shortestSuperstring(A))
}

func shortestSuperstring(A []string) string {
	m := make(map[string]map[string]int)
	for i := 0; i < len(A); i++ {
		m[A[i]] = make(map[string]int)
		for j := 0; j < len(A); j++ {
			if i != j {
				v := checkSame(A[i], A[j])
				if v > 0 {
					m[A[i]][A[j]] = v
				}
			}
		}
	}
	ret := ""
	for i := 0; i < len(A); i++ {
		ret += A[i]
	}
	if len(m) == 0 {
		return ret
	}
//	for k := range m {
//		used := make(map[string]bool)
//		used[k] = true
//		r := dfs(A, m, k, used, k)
//		fmt.Println(r)
//		fmt.Println(used)
//		used[k] = false
//		if len(r) < len(ret) {
//			ret = r
//		}
//	}
	return dfs(A, m, "", make(map[string]bool), "")
}

func dfs(A []string, m map[string]map[string]int, k string, used map[string]bool, ret string) string {
	r := ""
	if _, ok := m[k]; ok {
		for k1, v1 := range m[k] {
			if !used[k1] {
				used[k1] = true
				t := k1[v1:]
				t = dfs(A, m, k1, used, t)
				used[k1] = false
				if r == "" || len(r) > len(t) {
					r = t
				}
			}
		}
	} else {
		flag := false
		for k1 := range m {
			if !used[k1] {
				flag = true
				used[k1] = true
				t := k1
				t = dfs(A, m, k1, used, t)
				used[k1] = false
				if r == "" || len(r) > len(t) {
					r = t
				}
			}
		}
		if !flag {
			for i := 0; i < len(A); i++ {
				if !used[A[i]] {
					ret += A[i]
				}
			}
			return ret
		}
	}
	return ret + r
}

func checkSame(s, t string) int {
	last := len(s) - 1
	max := 0
	for j := 0; j < len(t); j++ {
		if t[j] == s[last] {
			k := j - 1
			i := last - 1
			for k >= 0 && i >= 0 && t[k] == s[i] {
				k--
				i--
			}
			if k < 0 && i >= 0 {
				if j + 1 > max {
					max = j + 1
				}
			}
		}
	}
	return max
}
