package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(decodeString("3[a3[ert]bb]2[bc]"))
}

func decodeString(s string) string {
	l := list.New()
	for i := 0; i < len(s); {
		if s[i] != ']' {
			nums := 0
			for s[i] >= '0' && s[i] <= '9' {
				d, _ := strconv.Atoi(string(s[i]))
				nums = nums*10 + d
				i++
			}
			if nums != 0 {
				l.PushBack(strconv.Itoa(nums))
			}
			l.PushBack(string(s[i]))
			i++
		} else {
			v := l.Back()
			v1 := v.Value.(string)
			str := ""
			mid := ""
			for v1 != "[" {
				str = v1 + str
				l.Remove(v)
				v = l.Back()
				v1 = v.Value.(string)
			}
			l.Remove(v)
			v = l.Back()
			v1 = v.Value.(string)
			num, _ := strconv.Atoi(v1)
			for i := 0; i < num; i++ {
				mid += str
			}
			l.Remove(v)
			l.PushBack(mid)

			i++
		}
	}
	ret := ""
	for e := l.Front(); e != nil; e = e.Next() {
		ret += e.Value.(string)
	}
	return ret
}
