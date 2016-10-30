package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	str := make([]string, n)

	for i := 0; i < n; i++ {
		num := i + 1
		if num%3 == 0  && num%5 == 0 {
			str[i] = "FizzBuzz"
		} else if num%3 == 0 {
			str[i] = "Fizz"
		} else if num%5 == 0 {
			str[i] = "Buzz"
		} else {
			str[i] = strconv.Itoa(num)
		}
	}

	return str
}
