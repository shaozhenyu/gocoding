package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//f := "K4(ON(SO3)2)2"
	f := "Ksad(Msa12T)3SF"
	//f := "T(A(BC)2)3"
	fmt.Println("origin: ", f)
	fmt.Println(countOfAtoms(f))
}

func handleBrackets(formula string, index int) (map[string]int, int) {
	m := make(map[string]int)
	if index >= len(formula) {
		return m, index
	}
	m, index1 := countOfAtoms1(formula[index:])
	index1 += index
	count, index2 := getSize(index1, formula)
	for key := range m {
		m[key] *= count
	}
	return m, index2
}

func countOfAtoms1(formula string) (map[string]int, int) {
	m := make(map[string]int)
	i := 0
	for i < len(formula) {
		if formula[i] == '(' {
			mBrackets, end := handleBrackets(formula, i+1)
			i = end
			for key := range mBrackets {
				m[key] += mBrackets[key]
			}
		} else if formula[i] == ')' {
			return m, i + 1
		} else {
			child, index := getEle(i+1, formula)
			key := string(formula[i]) + child
			count, index := getSize(index, formula)
			m[key] += count
			i = index
		}
	}
	return m, i + 1
}

func countOfAtoms(formula string) string {
	m, _ := countOfAtoms1(formula)
	fmt.Println(m)

	str := make([]string, len(m))
	index := 0
	for key := range m {
		str[index] = key
		index++
	}
	sort.Strings(str)
	ret := ""
	for i := 0; i < len(str); i++ {
		if m[str[i]] == 1 {
			ret += str[i]
		} else {
			ret += str[i] + strconv.Itoa(m[str[i]])
		}
	}
	return ret

}

func getEle(index int, formula string) (string, int) {
	e := ""
	for index < len(formula) && formula[index] >= 'a' && formula[index] <= 'z' {
		e += string(formula[index])
		index++
	}
	return e, index
}

func getSize(index int, formula string) (int, int) {
	start := index
	for index < len(formula) && formula[index] >= '0' && formula[index] <= '9' {
		index++
	}
	if start == index {
		return 1, index
	}
	val, _ := strconv.Atoi(formula[start:index])
	return val, index
}
