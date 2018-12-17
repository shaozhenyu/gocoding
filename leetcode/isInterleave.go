package main

import (
	"fmt"
	"container/list"
)

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbbaccc"

	// s1 = "aaba"
	// s2 = "aaa"
	// s3 = "aaabaaa"
	fmt.Println(isInterleave(s1, s2, s3))
}

type Node struct {
	i1, i2, i3 int
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) {
		return false
	}
	visited := make([][]bool, len(s1) + 1)
	for i := 0; i <= len(s1); i++ {
		visited[i] = make([]bool, len(s2) + 1)
	}
	l := list.New()
	l.PushBack(Node{0, 0, 0})
	for l.Len() > 0 {
		e := l.Front()
		v := e.Value.(Node)
		l.Remove(e)
		if v.i1 < len(s1) && s1[v.i1] == s3[v.i3] {
			i1 := v.i1 + 1
			i3 := v.i3 + 1
			if i3 == len(s3) {
				return true
			}
			if visited[i1][v.i2] {
				continue
			}
			visited[i1][v.i2] = true
			l.PushFront(Node{i1, v.i2, i3})
		}
		if v.i2 < len(s2) && s2[v.i2] == s3[v.i3] {
			i2 := v.i2 + 1
			i3 := v.i3 + 1
			if i3 == len(s3) {
				return true
			}
			if visited[v.i1][i2] {
				continue
			}
			visited[v.i1][i2] = true
			l.PushFront(Node{v.i1, i2, i3})
		}
	}
	return false
}
