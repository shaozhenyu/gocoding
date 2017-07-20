package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

// Teacher会优先调用本身的ShowA()，如果没有，会调用可以调用people的ShowA()
func main() {
	t := Teacher{}
	t.ShowB()
}
