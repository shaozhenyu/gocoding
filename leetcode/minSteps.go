package main

import "fmt"

func main() {
	for i := 2; i <= 15; i++ {
		fmt.Println(minSteps(i))
	}
}

//2
//3
//4
//5
//5
//7
//6
//6
//7
//11
//7
//13
//9
//8
//8
//17
//8
//19
//9

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			ret := i + minSteps(n/i)
			return ret
		}
	}
	return n
}
