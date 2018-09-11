package main

import "fmt"

type Trie struct {
	exist  bool
	childs map[byte]*Trie
}

func main() {
	c := Constructor()
	t := &c
	t.Insert("ap")
	fmt.Println(t.Search("ap"))
	fmt.Println(t.StartsWith("a"))
	fmt.Println(t.StartsWith("aa"))
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		childs: make(map[byte]*Trie, 26),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	t := this
	for i := 0; i < len(word); i++ {
		v := word[i]
		if _, ok := t.childs[v]; !ok {
			c := Constructor()
			t.childs[v] = &c
		}
		t = t.childs[v]
		if i == len(word)-1 {
			t.exist = true
		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	t := this
	for i := 0; i < len(word); i++ {
		v := word[i]
		if _, ok := t.childs[v]; !ok {
			return false
		}
		t = t.childs[v]
		if i == len(word)-1 && !t.exist {
			return false
		}
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for i := 0; i < len(prefix); i++ {
		v := prefix[i]
		if _, ok := t.childs[v]; !ok {
			return false
		}
		t = t.childs[v]
	}
	return true
}
