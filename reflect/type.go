package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = 10
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)
	fmt.Println("value : ", v)
	fmt.Println("type : ", t)

	//ta := v.Elem().Type() //panic
	//fmt.Println(ta)

	b := &a
	tb := reflect.ValueOf(b).Elem().Type()
	fmt.Println("type b : ", tb)

	m := map[string]interface{}{}

	m["a"] = "a"
	m["1"] = 1

	vm := reflect.ValueOf(m)
	tm := reflect.TypeOf(m)
	fmt.Println("value m : ", vm)
	fmt.Println("type m : ", tm)
	//fmt.Println("type m : ", vm.Elem().Type())
	fmt.Println("type m : ", vm.Elem())
	fmt.Println("m key type : ", tm.Key())
}
