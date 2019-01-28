package main

import "fmt"

func main() {
	A := []int{9,4,2,10,7,8,8,1,9}
	fmt.Println(maxTurbulenceSize(A))
}

func maxTurbulenceSize(A []int) int {
	if len(A) <= 1 {
		return len(A)
	}
	count := 1
	max := 1
	pre := A[0]
	flag := 0
	for i := 1; i < len(A); i++ {
		if A[i] == pre {
			flag = 0
			if count > max {
				max = count
			}
			count = 1
		} else if A[i] > pre {
			if flag <= 0 {
				flag = 1
				count++
			} else {
				if count > max {
					max = count
				}
				count = 2
			}
			pre = A[i]
		} else {
			if flag >= 0 {
				flag = -1
				count++
			} else {
				if count > max {
					max = count
				}
				count = 2
			}
			pre = A[i]
		}
	}
	if count > max {
		max = count
	}
	return max
}
