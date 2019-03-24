package main

import "fmt"

func main() {
	s := "AABABBA"
	k := 1
	fmt.Println(characterReplacement(s, k))
}

type Node struct {
	continues []int
	k int
	count int
}

func characterReplacement(s string, k int) int {
	n := [26]Node{}
	for i := 0; i < 26; i++ {
		n[i] = Node{make([]int, 1), k, 0}
	}
	max := 0
	for i := 0; i < len(s); i++ {
		c := int(s[i] - 'A')
		for j := 0; j < 26; j++ {
			if j != c {
				if n[j].k == 0 {
					n[j].count -= (n[j].continues[0] + 1)
					n[j].continues = n[j].continues[1:]
					n[j].k++
				}
				n[j].k--
				n[j].continues = append(n[j].continues, 0)
			} else {
				n[j].continues[len(n[j].continues)-1]++
			}
			n[j].count++
			if n[j].count > max {
				max = n[j].count
			}
		}
	}
	return max
}
