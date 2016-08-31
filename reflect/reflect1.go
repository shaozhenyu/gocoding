package main

import (
	"fmt"
	"reflect"
)

type Speaker interface {
	Speak() string
	Set(string)
}

type Teacher struct {
	Name string
}

func (this *Teacher) Speak() string {
	return this.Name
}

func (this *Teacher) Set(name string) {
	this.Name = name
}

func TestRef(s Speaker) (t reflect.Type) {
	reflectVal := reflect.ValueOf(s)
	t = reflectVal.Elem().Type()
	fmt.Printf("reflect.ValueOf(%v).Elem().Type() = %v\n", s, t)
	return
}

func TestRef2(s Speaker) (t reflect.Type) {
	t = reflect.TypeOf(s)
	fmt.Printf("reflect.TypeOf(%v)=%v\n", s, t)
	return
}

func TestRef3(s *Teacher) (t reflect.Type) {
	t = reflect.TypeOf(s).Elem()
	fmt.Printf("reflect.TypeOf(%v).Elem()=%v\n", s, t)
	return
}

func main() {
	te := &Teacher{
		Name: "szy",
	}
	fmt.Printf("source=%v\n", te)
	fmt.Println("--------")

	var tea Speaker = te

	fmt.Printf("tea value : %v\n", reflect.ValueOf(tea))
	fmt.Printf("te value : %v\n", reflect.ValueOf(te))
	fmt.Printf("tea type : %v\n", reflect.TypeOf(tea))
	fmt.Printf("te type : %v\n", reflect.TypeOf(te))
	fmt.Printf("\n\n\n")

	t1 := TestRef(tea)
	if m, ok := reflect.New(t1).Interface().(Speaker); ok {
		fmt.Printf("value of m : %v\n", reflect.ValueOf(m))
		fmt.Printf("type of m : %v\n", reflect.TypeOf(m))
		fmt.Printf("reflect.New(%v).Interface().(Speaker)=%v\n", t1, m)
		fmt.Printf("se.Speak()=%v\n", m.Speak())
		m.Set("lyy")
		fmt.Printf("reflect.New(%v).Interface().(Speaker)=%v\n", t1, m)
		fmt.Printf("se.Speak()=%v\n", m.Speak())
	}

	fmt.Printf("Source Instance te = %v\n", te)
	fmt.Printf("\n")

	t2 := TestRef2(tea)
	if m, ok := reflect.New(t2.Elem()).Interface().(Speaker); ok {
		fmt.Printf("reflect.New(%v.Elem()).Interface().Speaker=%v\n", t2, m)
		fmt.Printf("se.Speak()=%v\n", m.Speak())
		m.Set("anan")
		fmt.Printf("reflect.New(%v.Elem()).Interface().Speaker=%v\n", t2, m)
		fmt.Printf("se.Speak()=%v\n", m.Speak())
	}
	fmt.Printf("Source Instance te = %v\n", te)
	fmt.Printf("\n")

	t3 := TestRef3(te)
	if m, ok := reflect.New(t3).Interface().(Speaker); ok {
		fmt.Printf("reflect.New(%v).Interface().(Speaker)=%v\n", t3, m)
		fmt.Printf("se.Speak()=%v\n", m.Speak())
		m.Set("tibbers")
		fmt.Printf("reflect.New(%v).Interface().(Speaker)=%v\n", t3, m)
		fmt.Printf("se.Speak()=%v\n", m.Speak())
	}
	fmt.Printf("Source Instance te = %v\n", te)
	fmt.Printf("\n\n\n")

	fmt.Printf("reflect.TypeOf(tea)=%v\n", reflect.TypeOf(tea))
	fmt.Printf("reflect.TypeOf(tea).Elem()=%v\n", reflect.TypeOf(tea).Elem())
	fmt.Printf("reflect.ValueOf(tea)=%v\n", reflect.ValueOf(tea))
	fmt.Printf("reflect.ValueOf(tea).Elem()=%v\n", reflect.ValueOf(tea).Elem())
	fmt.Printf("\n\n\n")

	if se, ok := reflect.ValueOf(tea).Interface().(Speaker); ok {
		fmt.Printf("reflect.ValueOf(%v).Interface().(Speaker)=%v\n", tea, se)
		fmt.Printf("se.Speak()=%v\n", se.Speak())
		se.Set("aaa")
		fmt.Printf("reflect.ValueOf(%v).Interface().(Speaker)=%v\n", tea, se)
		fmt.Printf("se.Speak()=%v\n", se.Speak())
	}
	fmt.Printf("Source Instance te = %v\n", te)
	fmt.Printf("\n\n\n")

	if se, ok := reflect.New(reflect.TypeOf(tea).Elem()).Interface().(Speaker); ok {
		fmt.Printf("reflect.New(reflect.TypeOf(%v).Elem()).Interface().(Speaker)=%v\n", tea, se)
		fmt.Printf("se.Speak()=%v\n", se.Speak())
		se.Set("bbb")
		fmt.Printf("reflect.New(reflect.TypeOf(%v).Elem()).Interface().(Speaker)=%v\n", tea, se)
		fmt.Printf("se.Speak()=%v\n", se.Speak())
	}
	fmt.Printf("Source Instance te=%v\n", te)
	fmt.Printf("\n\n\n")

	if se, ok := reflect.New(reflect.TypeOf(te).Elem()).Interface().(Speaker); ok {
		fmt.Printf("reflect.New(reflect.TypeOf(%v).Elem()).Interface().(Speaker)=%v\n", te, se)
		fmt.Printf("se.Speak()=%v\n", se.Speak())
		se.Set("ccc")
		fmt.Printf("reflect.New(reflect.TypeOf(%v).Elem()).Interface().(Speaker)=%v\n", te, se)
		fmt.Printf("se.Speak()=%v\n", se.Speak())
	}
	fmt.Printf("Source Instance te=%v\n", te)
	fmt.Printf("\n\n\n")



}
