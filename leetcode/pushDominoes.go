package main

import "fmt"

func main() {
	d := ".L.R...LR..L.."
	fmt.Println(pushDominoes(d))
	fmt.Println("true : LL.RR.LLRRLL..")
}

func pushDominoes(dominoes string) string {
	if len(dominoes) <= 1 {
		return dominoes
	}
	str := []byte(dominoes)
	if dominoes[1] == 'L' && dominoes[0] != 'R' {
		str[0] = 'L'
	}
	if dominoes[len(dominoes)-2] == 'R' && dominoes[len(dominoes)-1] != 'L' {
		str[len(dominoes)-1] = 'R'
	}
	for i := 1; i < len(dominoes) - 1; i++ {
		if dominoes[i] != '.'  || (dominoes[i-1] == 'R' && dominoes[i+1] == 'L'){
			continue
		}
		if dominoes[i-1] == 'R' {
			str[i] = 'R'
		} else if dominoes[i+1] == 'L' {
			str[i] = 'L'
		}
	}
	if string(str) == dominoes {
		return string(str)
	}
	return pushDominoes(string(str))
}
