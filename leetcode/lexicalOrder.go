package main

import "fmt"

func main() {
	fmt.Println(lexicalOrder(200))
}

func lexicalOrder(n int) []int {
	ret := make([]int, n)
	idx := 0
	var lorder func(now int)
	lorder = func(now int) {
		if now > n {
			return
		}
		ret[idx] = now
		idx += 1
		if now * 10 <= n {
			lorder(now*10)
		}
		if now%10 == 0 {
			for i := 1; i <= 9; i++ {
				lorder(now+i)
			}
		}
	}
	for i := 1; i <= 9; i++ {
		lorder(i)
	}
	return ret
}
