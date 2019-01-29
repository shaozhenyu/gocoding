package main

import (
	"fmt"
)

func main() {
	people := []int{3,2,2,1}
	limit := 3
	fmt.Println(numRescueBoats(people, limit))
}

func numRescueBoats(people []int, limit int) int {
	count := 0
	bucket := make([]int, 30001)
	for i := 0; i < len(people); i++ {
		bucket[people[i]]++
	}
	for i := 1; i < len(bucket); i++ {
		if bucket[i] == 0 {
			continue
		}
		left := limit - i
		t := bucket[i]
		bucket[i] = 0
		for j := len(bucket) - 1; j > i; j-- {
			if j > left || bucket[j] == 0 {
				continue
			}
			if bucket[j] >= t {
				count += t
				bucket[j] -= t
				t = 0
				break
			} else {
				count += bucket[j]
				t -= bucket[j]
				bucket[j] = 0
			}
		}
		if t > 0 {
			if i + i <= limit {
				count += t/2 + t%2
			} else {
				count += t
			}
		}
	}
	return count
}
