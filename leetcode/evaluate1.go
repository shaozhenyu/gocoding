package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evaluate("(add 1 2)"))
	fmt.Println(evaluate("(let x 2 (mult x (let x 3 y 4 (add x y))))"))
}

func evaluate(expression string) int {
}

func cal(v)
