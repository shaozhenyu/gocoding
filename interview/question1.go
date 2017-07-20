package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	pase_student()
}

// stu这个变量在循环的时候地址不会变
func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
		fmt.Println(m)
	}
}
