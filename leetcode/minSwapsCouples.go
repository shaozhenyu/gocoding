package main

import "fmt"

func main() {
	row := []int{2,0,3,6,4,5,1,7}
	fmt.Println(minSwapsCouples(row))
}

func minSwapsCouples(row []int) int {
	fmt.Println("11: ", row)
	m := make(map[int]int)
	for i := 0; i < len(row); i++ {
		m[row[i]] = i
	}
	count := 0
	couple := 0
	for i := 0; i < len(row); i += 2 {
		if row[i]%2 == 0 {
			couple = row[i] + 1
		} else {
			couple = row[i] - 1
		}
		if row[i+1] != couple {
			coupleIndex := m[couple]
			m[row[i+1]] = coupleIndex
			m[couple] = i+1
			row[i+1], row[coupleIndex] = row[coupleIndex], row[i+1]
			count++
		}
	}
	fmt.Println("22: ", row)
	return count
}
