package main

import "fmt"

func main() {
	s := Constructor([]string{"abc", "f", "ff"})
	fmt.Println(s.Query('a'))
	fmt.Println(s.Query('f'))
	fmt.Println(s.Query('f'))
	fmt.Println(s.Query('a'))
	fmt.Println(s.Query('b'))
	fmt.Println(s.Query('c'))
}

func printS(s *Trie) {
	fmt.Println(s)
	for k, v := range s.childs {
		fmt.Println(string(k))
		printS(v)
	}
}

type StreamChecker struct {
	Trie *Trie
	LastTries []*Trie
}

type Trie struct {
	isEnd bool
	childs [26]*Trie
}

func Constructor(words []string) StreamChecker {
	exists := make(map[string]bool)
	trie := &Trie{
		childs: [26]*Trie{},
	}
	for i := 0; i < len(words); i++ {
		if exists[words[i]] {
			continue
		}
		trie.Insert(words[i])
		exists[words[i]] = true
	}
	return StreamChecker{
		Trie: trie,
		LastTries: make([]*Trie, 0),
	}
}


func (this *StreamChecker) Query(letter byte) bool {
	ret := false
	newTries := make([]*Trie, 0, 1000)
	lett := int(letter - 'a')
	for i := 0; i < len(this.LastTries); i++ {
		if this.LastTries[i].childs[lett] != nil {
			child := this.LastTries[i].childs[lett]
			if child.isEnd == true {
				ret = true
			}
			newTries = append(newTries, child)
		}
	}
	if this.Trie.childs[lett] != nil {
		child := this.Trie.childs[lett]
		newTries = append(newTries, child)
		if child.isEnd == true {
			ret = true
		}
	}
	this.LastTries = newTries
	return ret
}

func newTrie() *Trie {
	return &Trie{
		childs: [26]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	t := this
	for i := 0; i < len(word); i++ {
		v := int(word[i] - 'a')
		if t.childs[v] == nil {
			c := newTrie()
			t.childs[v] = c
		}
		t = t.childs[v]
		if i == len(word)-1 {
			t.isEnd = true
		}
	}
}
