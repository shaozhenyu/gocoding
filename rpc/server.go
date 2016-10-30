package main

import (
	"log"
	"fmt"

	"net"
	"net/http"
	"net/rpc"
)

type Echo int

type Nums struct {
	A, B, C int
}

func (e *Echo) Hi(args string, reply *string) error {
	*reply = "echo : " + args
	return nil
}

func (e *Echo) Sum(args int, reply *int) error {
	*reply = args + args
	return nil
}

func (e *Echo) Add(args Nums, reply *int) error {
	fmt.Println(args.A, args.B, args.C)

	*reply = args.A + args.B
	return nil
}

func main() {
	rpc.Register(new(Echo))
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("listen error : ", err)
	}

	http.Serve(l, nil)
}
