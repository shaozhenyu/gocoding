package main

import (
	//"container/list"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(diffWaysToCompute("1+2+3+4"))
}

func diffWaysToCompute(input string) []int {
	ret := []int{}
	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			p1 := input[:i]
			p2 := input[i+1:]
			s1 := diffWaysToCompute(p1)
			s2 := diffWaysToCompute(p2)
			for a := 0; a < len(s1); a++ {
				for b := 0; b < len(s2); b++ {
					switch input[i] {
					case '+':
						ret = append(ret, s1[a]+s2[b])
					case '-':
						ret = append(ret, s1[a]-s2[b])
					case '*':
						ret = append(ret, s1[a]*s2[b])
					}
				}
			}
		}
	}
	if len(ret) == 0 {
		num, _ := strconv.Atoi(string(input))
		ret = append(ret, num)
	}
	return ret
}

// func diffWaysToCompute(input string) []int {
// 	l := list.New()

// 	num := ""
// 	for i := 0; i < len(input); i++ {
// 		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
// 			n, _ := strconv.Atoi(num)
// 			l.PushBack(n)
// 			l.PushBack(string(input[i]))
// 			num = ""
// 		} else {
// 			num += string(input[i])
// 		}
// 	}
// 	n, _ := strconv.Atoi(num)
// 	l.PushBack(n)

// 	doNext(l)

// 	return []int{}
// }

// func print(key string, l *list.List) {
// 	fmt.Printf("%s", key)
// 	for e := l.Front(); e != nil; e = e.Next() {
// 		fmt.Printf("%v", e.Value)
// 	}
// 	fmt.Println()
// }

// func doNext(l *list.List) {
// 	if l.Len() == 1 {
// 		fmt.Println(l.Front().Value)
// 	}

// 	for e := l.Front(); e != nil; e = e.Next() {
// 		switch e.Value {
// 		case "+":
// 			print("start :", l)
// 			pre := l.Remove(e.Prev())
// 			next := l.Remove(e.Next())
// 			value := e.Value
// 			e.Value = pre.(int) + next.(int)
// 			newL := list.New()
// 			newL = l
// 			e.Value = value
// 			l.InsertBefore(pre, e)
// 			l.InsertAfter(next, e)
// 			print("midd1: ", l)
// 			print("new:", newL)
// 			doNext(newL)
// 			print("midd2: ", l)
// 			print("end :", l)
// 		case "-":
// 			print("start :", l)
// 			pre := l.Remove(e.Prev())
// 			next := l.Remove(e.Next())
// 			value := e.Value
// 			e.Value = pre.(int) - next.(int)
// 			doNext(l)
// 			e.Value = value
// 			l.InsertBefore(pre, e)
// 			l.InsertAfter(next, e)
// 			print("end :", l)
// 		case "*":
// 			print("start :", l)
// 			pre := l.Remove(e.Prev())
// 			next := l.Remove(e.Next())
// 			value := e.Value
// 			e.Value = pre.(int) * next.(int)
// 			doNext(l)
// 			e.Value = value
// 			l.InsertBefore(pre, e)
// 			l.InsertAfter(next, e)
// 			print("end :", l)
// 		default:
// 		}
// 	}

// }
