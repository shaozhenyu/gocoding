package main

import "fmt"

func main() {
	// fmt.Println(brokenCalc(3, 10))
	// fmt.Println(brokenCalc(5, 8))
	fmt.Println(brokenCalc(3, 13))
}

func brokenCalc(X int, Y int) int {
	count := 0
	for X < Y {
		count += 1 + Y%2
		Y = (Y + Y%2)/2
	}
	return count + X - Y
}
