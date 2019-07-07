package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(digitsCount(3, 100, 250))
	fmt.Println(digitsCount(0, 1, 10))
	fmt.Println(digitsCount(0, 1, 20))
}

func digitsCount(d int, low int, high int) int {
	add := 0
	t := low
	for t > 0 {
		if t%10 == d {
			add++
		}
		t = t/10
	}
	ret := dc(d, high) - dc(d, low) + add
	if d == 0 {
		ret -= 10
	}
	return ret
}


func dc(d int, n int) int {
	ret := 0
	low, high := 1, n
	i := 1
	for high != 0 {
		high = n/int(math.Pow(10, float64(i)))
		tmp := n%int(math.Pow(10, float64(i)))
		low =  tmp%int(math.Pow(10, float64(i-1)))
		now := tmp/int(math.Pow(10, float64(i-1)))
		if now == d {
			ret += high * int(math.Pow(10, float64(i-1))) + low + 1
		} else if now < d {
			ret += high * int(math.Pow(10, float64(i-1)))
		} else {
			ret += (high+1) *int(math.Pow(10, float64(i-1)))
		}
		i++
	}
	if d == 0 {
	//	ret /= 10
	}
	fmt.Println(d, n,ret)
	return ret
}

func dc1(d int, n int) int {
	start := d
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
	}
	return count
}


