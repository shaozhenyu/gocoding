package main

import "fmt"

func main() {
	p1 := []int{0,0}
	p2 := []int{1,1}
	p3 := []int{1,0}
	p4 := []int{0,1}
	fmt.Println(validSquare(p1, p2, p3, p4))
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	p12 := square(p1, p2)
	p13 := square(p1, p3)
	p14 := square(p1, p4)
	p23 := square(p2, p3)
	p24 := square(p2, p4)
	p34 := square(p3, p4)
	if p12 == 0 || p13 == 0 || p14 == 0 || p23 == 0 || p24 == 0 || p34 == 0 {
		return false
	}
	if p12 == p13 && (p12 + p13) == p14 {
		if p24 == p34 && p24 == p12 && p23 == p14 {
			return true
		}
	}
	if p12 == p14 && (p12 + p14) == p13 {
		if p23 == p34 && p23 == p12 && p24 == p13 {
			return true
		}
	}
	if p13 == p14 && (p13 + p14) == p12 {
		if p23 == p24 && p24 == p14 && p12 == p34 {
			return true
		}
	}
	return false
}

func square(a , b []int) int {
	return (a[0] - b[0]) * (a[0] - b[0]) + (a[1] - b[1]) * (a[1] - b[1])
}
