package main

import "fmt"

func main() {
	S := "abc"
	shifts := []int{3,5,9}
	fmt.Println(shiftingLetters(S, shifts))
}

var low = []byte{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}

func shiftingLetters(S string, shifts []int) string {
	bs := []byte(S)
	move := 0
	for i := len(bs)-1; i >= 0; i-- {
		move += shifts[i]
		bs[i] = shift(bs[i], move)
	}
	return string(bs)
}

func shift(b byte, move int) byte {
	return low[(int(b - 'a') + move)%26]
}
