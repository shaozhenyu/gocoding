package main

import "fmt"
import "list/singlechain"

func main() {
	fmt.Println("singlechain example")
	var h singlechain.Node
	for i := 1; i <= 10; i++ {
		var d singlechain.Node
		d.Data = i
		h.Insert(&d, i)
	}

	pHead := h.GetHead()
	fmt.Println("h head : ", pHead.Data)
	pLast := h.GetLast()
	fmt.Println("h last : ", pLast.Data)
	
	var d singlechain.Node
	d.Data = 100
	h.InsertOnLast(&d)

	fmt.Println("h List : ")
	h.List()

	length := h.GetLen()
	fmt.Println("h len : ", length)
}

