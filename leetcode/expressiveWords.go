package main

import "fmt"

func main() {
	S := "heeellooo"
	words := []string{"hello", "hi", "helo"}
	fmt.Println(expressiveWords(S, words))
}

type Node struct {
	val        byte
	count      int
	isExtended bool
}

func expressiveWords(S string, words []string) int {
	if len(S) == 0 {
		return 0
	}
	pre := S[0]
	count := 1
	nodes := make([]Node, 0)
	for i := 1; i < len(S); i++ {
		if S[i] != pre {
			node := Node{pre, count, false}
			if count >= 3 {
				node.isExtended = true
			}
			nodes = append(nodes, Node{pre, count, count >= 3})
			count = 0
		}
		count++
		pre = S[i]
	}
	nodes = append(nodes, Node{pre, count, count >= 3})
	// fmt.Println(nodes)
	ret := 0
	for i := 0; i < len(words); i++ {
		word := words[i]
		k := 0
		if len(word) == 0 {
			continue
		}
		pre := word[0]
		count := 1
		j := 1
		for ; j < len(word); j++ {
			// fmt.Println("qq:", count, string(word[j]))
			if word[j] != pre {
				if k < len(nodes) && nodes[k].val == pre && (nodes[k].isExtended && count <= nodes[k].count || count == nodes[k].count) {
					k++
					count = 0
				} else {
					// fmt.Println("xxxxx", j, k, nodes[k], count)
					break
				}
			}
			count++
			pre = word[j]
		}
		// fmt.Println(word, j, k)
		if j == len(word) && k == len(nodes)-1 && nodes[k].val == pre && (nodes[k].isExtended && count <= nodes[k].count || count == nodes[k].count) {
			ret++
		}
	}
	return ret
}
