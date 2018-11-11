package main

import "fmt"

func main() {
	A := []int{5,4,0,3,1,6,2}
	fmt.Println(arrayNesting(A))
}

func arrayNesting(nums []int) int {
    idxM := make(map[int]int)
	nM := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		idxM[i] = nums[i]
		nM[nums[i]] = i
	}
	max := 0
	for k, v := range idxM {
		count := 1
		delete(idxM, k)
		t := v
		for {
			if _, ok := idxM[t]; !ok {
				break
			}
			tmp := idxM[t]
			delete(idxM, t)
			delete(nM, tmp)
			t = tmp
			count++
		}
		t = k
		delete(nM, k)
		for {
			if _, ok := nM[t]; !ok {
				break
			}
			tmp := nM[t]
			delete(nM, t)
			delete(idxM, tmp)
			t = tmp
			count++
		}
		if count > max {
			max = count
		}
	}
	return max
}
