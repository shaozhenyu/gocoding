package main

import "fmt"

func main() {
	fmt.Println(longestArithSeqLength([]int{44,46,22,68,45,66,43,9,37,30,50,67,32,47,44,11,15,4,11,6,20,64,54,54,61,63,23,43,3,12,51,61,16,57,14,12,55,17,18,25,19,28,45,56,29,39,52,8,1,21,17,21,23,70,51,61,21,52,25,28}))
}

func longestArithSeqLength(A []int) int {
	exist := make(map[int][]int) // map[val]idx
	same := make(map[int]int)
	max := -1
	for i := 0; i < len(A); i++ {
		if _, ok := exist[A[i]]; !ok {
			exist[A[i]] = make([]int, 0)
		}
		exist[A[i]] = append(exist[A[i]], i)
		same[A[i]]++
		if same[A[i]] > max {
			max  = same[A[i]]
		}
	}
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			d := A[j] - A[i]
			if d == 0 {
				continue
			}
			preVal := A[j] + d
			preIdx := j
			count := 2
			for {
				if idxs, ok := exist[preVal]; ok {
					idx := -1
					for x := 0; x < len(idxs); x++ {
						if idxs[x] > preIdx {
							idx = idxs[x]
							break
						}
					}
					if idx == -1 {
						break
					}
					count++
					preVal += d
					preIdx = idx
				} else {
					break
				}
			}
			if count > max {
				max = count
			}
		}
	}
	return max
}
