package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name     string
	Age      int64
	Birthday string
}

func main() {
	user := &User{}

	t := reflect.TypeOf(*user)
	//can set
	v := reflect.ValueOf(user).Elem()
	fmt.Println(v.NumField())
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s -- %v \n", t.Filed(k).Name, v.Field(k).Interface())
	}
}
