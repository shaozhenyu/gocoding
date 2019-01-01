package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	persons := []int{0, 1, 1, 0, 0, 1, 0}
	times := []int{0, 5, 10, 15, 20, 25, 30}
	t := Constructor(persons, times)
	fmt.Println(t.Q(15))
}

type TopVotedCandidate struct {
	winner []Winner
}

type Node struct {
	person int
	time   int
}

type Winner struct {
	person int
	times  int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	l := list.New()
	winners := make([]Winner, 0, len(times))
	pVote := make(map[int]*list.Element)
	for i := 0; i < len(persons); i++ {
		if l.Len() == 0 {
			pVote[persons[i]] = l.PushBack(Node{persons[i], 1})
		} else {
			next := l.Front()
			node := Node{persons[i], 0}
			e := pVote[persons[i]]
			if e != nil {
				node = e.Value.(Node)
				next = e.Next()
				l.Remove(e)
			}
			node.time += 1
			fmt.Println("time: ", times[i], node)
			for ; next != nil; next = next.Next() {
				nextNode := next.Value.(Node)
				if node.time < nextNode.time {
					pVote[persons[i]] = l.InsertBefore(node, next)
					break
				}
			}
			if next == nil {
				pVote[persons[i]] = l.PushBack(node)
			}
		}

		e := l.Back()
		winner := e.Value.(Node)
		winners = append(winners, Winner{winner.person, times[i]})
	}
	return TopVotedCandidate{winners}
}

func (this *TopVotedCandidate) Q(t int) int {
	fmt.Println(this.winner)
	idx := sort.Search(len(this.winner), func(i int) bool {
		return t < this.winner[i].times
	})
	idx -= 1
	if idx < 0 {
		idx = 0
	}
	return this.winner[idx].person
}
