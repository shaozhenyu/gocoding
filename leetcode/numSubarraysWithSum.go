package main

import "fmt"

func main() {
	A := []int{0,0,1,0,0}
	A = []int{0,0,0,0,1,0,0,0,0,0,1,1,0,0,0,0,0,1,0,0}
	S := 3
	fmt.Println(numSubarraysWithSum(A, S))
}

func numSubarraysWithSum(A []int, S int) int {
	if S == 0 {
		sum := 0
		count := 0
		for i := 0; i < len(A); i++ {
			if A[i] == 1 {
				count += (1 + sum) * sum/2
				sum = 0
			} else {
				sum++
			}
		}
		return count + (1 + sum) * sum/2
	}
	start, end := 0, 0
	sum, count := 0, 0
	for end < len(A) {
		if A[end] == 1 {
			sum += 1
		}
		fmt.Println(end, sum, S)
		if sum == S {
			count++
			end++
			endCount := 0
			for end < len(A) && A[end] == 0 {
				endCount++
				end++
			}
			count += endCount
			for start < end && A[start] == 0 {
				count += endCount + 1
				start++
			}
			start++
			sum--
		} else {
			end++
		}
	}
	return count
}
