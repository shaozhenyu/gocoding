package main

import "fmt"

func main() {
	A := []int{2,1,2,1,2,1,1,1,1,1,4,3}
	fmt.Println(numSubarrayBoundedMax(A, 2, 3))
}

func numSubarrayBoundedMax(A []int, L int, R int) int {
	count := 0
	small := 0
	for i := 0; i < len(A); i++ {
		if A[i] >= L && A[i] <= R {
			j := i + 1
			count += 1 + small
			for j < len(A) && A[j] <= R {
				count += 1 + small
				j++
			}
			small = 0
		} else if A[i] < L {
			small++
		} else if A[i] > R {
			small = 0
		}
	}
	return count
}
