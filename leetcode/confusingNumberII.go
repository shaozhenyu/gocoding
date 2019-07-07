package main

import "fmt"

func main() {
	fmt.Println(confusingNumberII(20))
}

var tas = map[int]int{6:9,9:6,0:0,1:1,8:8}

func confusingNumberII(N int) int {
	m := []int{0, 1, 6, 8, 9}
	okD := make([]int, 0, 1000)
	okD = append(okD, 0)
	ret := 0
	k := 1
	for k <= N {
		lenD := len(okD)
		for i := 0; i < lenD; i++ {
			for j := 1; j < len(m); j++ {
				now := k * m[j] + okD[i]
				if now <= N && check(now) {
					ret++
				}
				okD = append(okD, now)
			}
		}
		k *= 10
	}
	return ret
}

func check(x int) bool {
	old, nw := x, 0
	for x > 0 {
		t := x%10
		t = tas[t]
		nw = nw * 10 + t
		x /= 10
	}
	if nw == old {
		return false
	}
	fmt.Println("xx:", old, nw)
	return true
}
