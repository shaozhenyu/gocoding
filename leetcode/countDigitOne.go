package main

import "fmt"

func main() {
	// fmt.Println(countDigitOne(13), 6)
	// fmt.Println(countDigitOne(111), 36)
	// fmt.Println(countDigitOne(112), 38)
	fmt.Println(countDigitOne(100), 21)
	fmt.Println(countDigitOne(101), 23)
}

func countDigitOne(n int) int {
	start := 1
	count := 0
	for n >= start {
		count += n/(start*10) * start
		left := n%(start*10)
		if left < start {
			start *= 10
			continue
		}
		if left >= (start * 2) {
			count += start
		} else {
			count += left - start + 1
		}
		start *= 10
		fmt.Println(count)
	}
	return count
}
