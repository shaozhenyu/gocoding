package main

import (
	"container/list"
	"fmt"
)

func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	//deadends := []string{}
	target := "0202"
	fmt.Println(openLock(deadends, target))
}

func openLock(deadends []string, target string) int {
	deadendsMap := make(map[string]int)
	for i := 0; i < len(deadends); i++ {
		deadendsMap[deadends[i]] = 1
	}
	init := "0000"
	if deadendsMap[init] == 1 {
		return -1
	}
	if target == init {
		return 0
	}
	visited := make(map[string]int)

	l := list.New()
	l.PushBack(target)

	var size int
	var str string
	step := 0
	for l.Len() > 0 {
		size = l.Len()
		for size > 0 {
			v := l.Front()
			str = v.Value.(string)
			l.Remove(v)
			size--

			if isOk(str) {
				fmt.Println("ok")
				return step
			}

			if visited[str] != 0 {
				continue
			}

			visited[str] = 1

			for i := 0; i < 4; i++ {
				str1 := str[:i] + string(((str[i]-48)+1+10)%10+48) + str[i+1:]
				if visited[str1] == 0 && deadendsMap[str1] == 0 {
					if isOk(str1) {
						return step + 1
					}
					l.PushBack(str1)
				}

				str2 := str[:i] + string(((str[i]-48)-1+10)%10+48) + str[i+1:]
				if visited[str2] == 0 && deadendsMap[str2] == 0 {
					if isOk(str2) {
						return step + 1
					}
					l.PushBack(str2)
				}
			}
		}
		step++
	}
	return step
}

// func search(m map[string]int, target string, v map[string]int) int {
// 	fmt.Println(t)
// 	if step >= min {
// 		fmt.Println("000", min)
// 		return min
// 	}
// 	if isOk(t) {
// 		fmt.Println("ok: ", t, step)
// 		return step
// 	}
// 	if m[pStr(t)] == 1 {
// 		fmt.Println("aaa")
// 		return -1
// 	}

// 	method := [][]int{{0, 1}, {0, -1}, {1, 1}, {1, -1}, {2, 1}, {2, -1}, {3, 1}, {3, -1}}
// 	var tmp1, tmp2 int
// 	tmp2 = 2 << 31
// 	for i := 0; i < len(method); i++ {
// 		a, b := method[i][0], method[i][1]
// 		if v[a][(t[a]+b+10)%10] == 0 {
// 			t[a] = (t[a] + b + 10) % 10
// 			v[a][t[a]] = 1
// 			tmp1 = search(m, t, v, step+1, tmp2)
// 			v[a][t[a]] = 0
// 			t[a] = (t[a] - b + 10) % 10

// 			fmt.Println("tmp1:", tmp1, tmp2)
// 			if tmp1 > 0 && tmp1 < tmp2 {
// 				tmp2 = tmp1
// 			}
// 		}
// 	}

// 	// fmt.Println("tmp2 step:", tmp2, step)
// 	// if tmp2 > 0 && tmp2 < step {
// 	// 	step = tmp2
// 	// }

// 	return tmp2
// }

func isOk(t string) bool {
	if t == "0000" {
		return true
	}
	return false
}
