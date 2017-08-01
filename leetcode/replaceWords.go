package main

import (
	"fmt"
	"strings"
)

func main() {
	dict := []string{"a", "aa", "aaa", "aaaa"}
	sentence := "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"

	replaceWords(dict, sentence)
}

func replaceWords(dist []string, sentence string) {
	root := New("")
	fmt.Println(root)
	buildTrie(dist, root)

	s := strings.Split(sentence, " ")
	for i := 0; i < len(s); i++ {
		one := s[i]
		temp := root
		for j := 0; j < len(one); j++ {
			this := temp.children[one[j]-'a']
			if this == nil {
				break
			} else {
				if this.isWord == true {
					s[i] = one[:j+1]
					break
				} else {
					temp = this
				}
			}
		}
	}
}

func buildTrie(dist []string, root *Trie) {
	lenD := len(dist)
	for i := 0; i < lenD; i++ {
		this := root
		d := dist[i]
		for j := 0; j < len(d); j++ {
			if this.children[d[j]-'a'] == nil {
				this.children[d[j]-'a'] = New(string(d[j]))
			}
			this = this.children[d[j]-'a']
		}
		this.isWord = true
	}
}

func New(val string) *Trie {
	return &Trie{
		val:      val,
		isWord:   false,
		children: make([]*Trie, 26),
	}
}

type Trie struct {
	val      string
	isWord   bool
	children []*Trie
}

// func replaceWords(dict []string, sentence string) {
// 	s := strings.Split(sentence, " ")

// 	var short int
// 	for k, one := range s {
// 		lenD := len(dict)
// 		lenOne := len(one)
// 		short = lenOne
// 		for i := 0; i < lenD; i++ {
// 			if strings.HasPrefix(one, dict[i]) {
// 				if short < len(dict[i]) {
// 					dict = append(dict[:i], dict[i+1:]...)
// 					lenD--
// 				} else if short > len(dict[i]) {
// 					short = len(dict[i])
// 					s[k] = dict[i]
// 				}
// 			}
// 		}
// 	}

// 	fmt.Println(s)
// }
