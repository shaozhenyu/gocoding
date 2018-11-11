package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is", "is", "a", "a", "a", "a"}
	fmt.Println(topKFrequent(words, 4))
	fmt.Println(topKFrequent1(words, 4))
}

type Fre struct {
	Word  string
	Times int
}

type FreHeap []Fre

func (h FreHeap) Len() int {
	return len(h)
}

func (h FreHeap) Less(i, j int) bool {
	if h[i].Times == h[j].Times {
		return h[i].Word > h[j].Word
	}
	return h[i].Times < h[j].Times
}
func (h FreHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *FreHeap) Push(x interface{}) {
	*h = append(*h, x.(Fre))
}
func (h *FreHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}

	h := &FreHeap{}
	heap.Init(h)
	for key, v := range m {
		heap.Push(h, Fre{key, v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	i := k - 1
	ret := make([]string, k)
	for h.Len() != 0 {
		v := heap.Pop(h).(Fre)
		ret[i] = v.Word
		i--
	}
	return ret
}

func topKFrequent1(words []string, k int) []string {
	m := make(map[string]int)
	for _, v := range words {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	fre := make([]Fre, len(m))
	var i int = 0
	for k, v := range m {
		fre[i] = Fre{
			Word:  k,
			Times: v,
		}
		i++
	}
	sort.Slice(fre, func(i, j int) bool {
		if fre[i].Times == fre[j].Times {
			com := strings.Compare(fre[i].Word, fre[j].Word)
			if com > 0 {
				return false
			} else {
				return true
			}
		}
		return fre[i].Times > fre[j].Times
	})
	ret := make([]string, k)
	for i = 0; i < k; i++ {
		ret[i] = fre[i].Word
	}
	return ret
}
