package main

import (
	"fmt"
	"container/list"
)

func main() {
	start := "AAA"
	end := "TTT"
	bank := []string{"AAT", "TTT", "ATT", "GAT", "GTT"}
	fmt.Println(minMutation(start, end, bank))
}

type Node struct {
	v string
	step int
}

// "A", "C", "G", "T"
func minMutation(start string, end string, bank []string) int {
	change := []byte{'A', 'C', 'G', 'T'}
	m := make(map[string]bool)
	for i := 0; i < len(bank); i++ {
		m[bank[i]] = true
	}
	if len(m) == 0 || !m[end] {
		return -1
	}
	queue := list.New()
	queue.PushBack(Node{start, 1})
	exist := make(map[string]bool)
	for queue.Len() > 0 {
		e := queue.Front()
		now := e.Value.(Node)
		queue.Remove(e)
		fmt.Println(now)
		if exist[now.v] {
			continue
		}
		for i := 0; i < len(now.v); i++ {
			b := []byte(now.v)
			for j := 0; j < 4; j++ {
				if b[i] == change[j] {
					continue
				}
				b[i] = change[j]
				fmt.Println("aa:", string(b))
				if string(b) == end {
					return now.step
				}
				if !exist[string(b)] && m[string(b)] {
					queue.PushBack(Node{string(b), now.step+1})
				}
			}
		}
		exist[now.v] = true
	}
	return -1
}
