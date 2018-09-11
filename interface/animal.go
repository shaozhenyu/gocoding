package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Play()
}

type Cat struct {
}

type Dog struct {
}

func (d *Dog) Eat() {
	fmt.Println("dog eat")
}

func (d *Dog) Play() {
	fmt.Println("dog play")
}

func (d *Dog) Jiao() {
	fmt.Println("dog jiao")
}

func (c *Cat) Eat() {
	fmt.Println("cat eat")
}

func main() {
	a := &Dog{}
	eat(a)
}

func eat(a Animal) {
	a.Eat()
	a.Jiao()
}
