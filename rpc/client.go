package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Nums struct {
	A, B, C int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dailing:", err)
	}

	args := "hello rpc"
	var reply string

	err = client.Call("Echo.Hi", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Printf("%s -- %s\n", args, reply)

	args1 := &Nums{1, 2, 3}
	var reply1 int

	err = client.Call("Echo.Add", args1, &reply1)

	fmt.Println(reply1)
}
