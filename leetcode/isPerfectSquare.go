package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(100))
}

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	var val int
	for i := 2; i <= num/2; i++ {
		val = i * i
		if val < num {
			continue
		}

		if i * i == num {
			return true
		}
		return false
	}
	return false
}
