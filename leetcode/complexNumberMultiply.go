package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(complexNumberMultiply("1+1i", "1+1i"))
}


func complexNumberMultiply(a string, b string) string {
	ap, ae := split(a)
	bp, be := split(b)
	fmt.Println(ap, ae)
	fmt.Println(bp, be)
	p := ap * bp - ae * be
	e := ap * be + bp * ae
	fmt.Println(p, e)
	return fmt.Sprintf("%d+%di", p, e)
}

func split(str string) (p, e int) {
	str = str[:len(str)-1]
	s := strings.Split(str, "+")
	p, _ = strconv.Atoi(s[0])
	e, _ = strconv.Atoi(s[1])
	return
}
