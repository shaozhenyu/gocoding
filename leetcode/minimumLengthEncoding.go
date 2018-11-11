package main

import "fmt"

func main() {
	// words := []string{"tellme", "bell", "me"}
	words := []string{"time", "atime", "btime"}
	fmt.Println(minimumLengthEncoding(words))
}

func minimumLengthEncoding(words []string) int {
	w := Constructor()
	for i := 0; i < len(words); i++ {
		w.AddWord(words[i])
	}
	return w.Search(0)
}

func (this *WordDictionary) Search(pre int) int {
	count := 0
	for i := 0; i < 26; i++ {
		if this.exist[i] {
			fmt.Println(i, count, pre)
			count += this.child[i].Search(pre+1)
		}
	}
	if count == 0 {
		count = pre + 1
	}
	return count
}

type WordDictionary struct {
	isEnd bool
	exist [26]bool
	child [26]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		exist: [26]bool{},
		child: [26]*WordDictionary{},
	}
}

func initW() *WordDictionary {
	return &WordDictionary{
		exist: [26]bool{},
		child: [26]*WordDictionary{},
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		this.isEnd = true
		return
	}
	v := word[len(word)-1] - 'a'
	if !this.exist[v] {
		this.exist[v] = true
		this.child[v] = initW()
	}
	this.child[v].AddWord(word[:len(word)-1])
}

