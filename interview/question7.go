package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func main() {
	u := &UserAges{
		ages: make(map[string]int),
	}

	u.Add("aa", 12)
	u.Add("bb", 13)
	u.Add("cc", 14)
	u.Add("aa", 15)
	fmt.Println(u.Get("aa"),u.Get("bb"), u.Get("cc"))
	fmt.Println(u)
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
