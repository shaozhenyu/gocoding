package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string `json:"name" bson:"name"`
	Age int64 `json:"age" bson:"age"`
	Address string `json:""`
}

func (p *People) Hello() (int64, int64) {
	fmt.Println("hello ", p.Name)
	return 1, 2
}

func (p *People) Bye() int64 {
	fmt.Println("Bye ", p.Name)
	return 3
}

func main() {
	var a float64
	a = 10

	T := reflect.TypeOf(a)
	V := reflect.ValueOf(a)

	fmt.Println("type a : ", T, V.Kind())
	fmt.Println("value a : ",V)

	fmt.Println("type name : ", T.Name())
	fmt.Println("type kind : ", T.Kind())

	//kind()
	if T.Kind() == reflect.Float64 {
		fmt.Println("type is float64")
	}

	//Interface
	y := V.Interface().(float64) //if not float64, panic
	fmt.Println("y is : ", y)
	fmt.Println("y type : ", reflect.TypeOf(y))

	//convert
	//i := V.Convert(reflect.Kind().Int64)
	//fmt.Println("type i : ", i.Type)

	people := &People{"lyy", 20, "yixing"}
	peopleV := reflect.ValueOf(people).Elem()
	peopleT := peopleV.Type()
	for i := 0; i < peopleV.NumField(); i++ {
		field := peopleV.Field(i)
		fmt.Printf("%s -- %s --  %v\n", peopleT.Field(i).Name, field.Type(), field.Interface())
	}

	fmt.Println("--------------------------")
	//Struct Tag
	p := People{"szy", 18, "yixing"}
	pt := reflect.TypeOf(p)
	pv := reflect.ValueOf(p)
	fmt.Println("type p : ", pt)

	for i := 0; i < pt.NumField(); i++ {
		field := pt.Field(i)
		fmt.Println("json: ", field.Tag.Get("json"))
		if value, ok := field.Tag.Lookup("bson"); ok {
			fmt.Println("bson: ", value)
		} else {
			fmt.Println("field have no bson tag:", field.Name)
		}
	}
	
	fmt.Println("--------------------------")
	//NumMethod
	pl := new(People)
	pl.Name = "lyy"
	plt := reflect.TypeOf(pl)
	plv := reflect.ValueOf(pl)

	//diff in pl and p
	fmt.Println("pl type : ", plt)
	fmt.Println("type p : ", pt)

	fmt.Println("--------------------------")

	for i := 0; i < plt.NumMethod(); i++ {
		method := plt.Method(i)
		fmt.Println("method name : ", method.Name)
		fmt.Println("method Type : ", method.Type)
		fmt.Println("method index : ", method.Index)
		ret := plv.MethodByName(method.Name).Call([]reflect.Value{})
		fmt.Println(ret[0])
	}

	//New
	n := reflect.New(plv.Elem().Type()).Interface()
	fmt.Printf("new pl value : %v\n", reflect.ValueOf(n))
	fmt.Printf("new pl type : %v\n", reflect.TypeOf(n))

	//canset pl and p
	fmt.Println("--------------------------")
	fmt.Println(reflect.ValueOf(&pl.Name).Elem().CanSet())
	fmt.Println(pv.FieldByName("name").CanSet())

	fmt.Println("--------------------------")
	//canset elem
	s := "abc"
	sv := reflect.ValueOf(&s)
	fmt.Println(sv.Elem().CanSet())
	fmt.Println(sv.CanSet())
	sv.Elem().SetString("bbb")
	fmt.Println("set string s : ", s)
}
