package main

import "fmt"

func main() {
	ratings := []int{5,4,1,1,0,0,3,2,1}
	fmt.Println(candy(ratings))
}

func candy(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}
	count := len(ratings)
	pre := ratings[0]
	idx := 1
	for idx < len(ratings) {
		de := 0
		for idx < len(ratings) && ratings[idx] <= pre {
			if ratings[idx] == pre {
				idx++
				count += de
				continue
			}
			pre = ratings[idx]
			idx++
			de++
			count += de
		}
		if idx == len(ratings) {
			break
		}
		pre = ratings[idx]
	}
	return count
}
