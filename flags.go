package main

import (
	"flag"
	"fmt"
)

//type interval []time.Duration
//var intervalF interval
//func init() {
//	flag.Var(&intervalF, "intervalF", "my intervalF")
//}

//result return *string
//var name *string = flag.String("name", "szy", "please input name")
var name = flag.String("name", "szy", "my name")

var age int
func init() {
	flag.IntVar(&age, "age", 18, "my age")
}


func main() {
	flag.Parse()
	fmt.Println(*name, age)
	fmt.Println(*name, age)
	fmt.Println(*name, age)
}
