package main

import "fmt"

func main() {
	pushed := []int{8,2,1,4,7,9,0,3,5,6}
	popped := []int{1,2,7,3,6,4,0,9,5,8}
	pushed = []int{1,2,3,4,5}
	popped = []int{4,5,3,2,1}
	fmt.Println(validateStackSequences(pushed, popped))
}

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}
	truePush := make([]int, len(pushed))
	idx := 0
	j := 0
	for i := 0; i < len(popped); i++ {
		if idx == 0 || truePush[idx-1] != popped[i] {
			for ; j < len(pushed) && pushed[j] != popped[i]; j++ {
				truePush[idx] = pushed[j]
				idx++
			}
			j++
		} else {
			idx--
		}
		fmt.Println(i,j,idx, truePush[:idx])
	}
	if idx == 0 {
		return true
	}
	return false
}
