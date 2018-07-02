package main

import (
	"fmt"
	"log"
	"reflect"
)

type User struct {
	Name     string
	Age      int64
	Birthday string
}

func (u *User) Add(args int64, reply *int64) error {
	*reply = args + u.Age
	return nil
}

func (u *User) Reduce(args int64, reply *int64) error {
	*reply = args - u.Age
	return nil
}

func main() {
	user := &User{}
	register(user)
}

func register(rcvr interface{}) {
	typ := reflect.TypeOf(rcvr)
	fmt.Println("type:", typ)
	val := reflect.ValueOf(rcvr)
	fmt.Println("value: ", val)
	fmt.Println("value name", reflect.Indirect(val).Type().Name())

	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mname := method.Name
		mtype := method.Type
		fmt.Println("method: ", mname, mtype)

		if method.PkgPath != "" {
			log.Fatal("pkg path should not be nil")
		}

		fmt.Println("numIn: ", mtype.NumIn())
	}
}
