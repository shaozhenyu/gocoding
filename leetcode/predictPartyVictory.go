package main

import "fmt"

func main() {
	s := "RDD"
	fmt.Println(predictPartyVictory(s))
}

func predictPartyVictory(senate string) string {
	return ppv([]byte(senate), 0, 0)
}

func ppv(senate []byte, r, d int) string {
	leftR := 0
	leftD := 0
	for i := 0; i < len(senate); i++ {
		if senate[i] == 'R' {
			if d == 0 {
				r++
				leftR++
			} else {
				senate[i] = ' '
				d--
			}
		} else if senate[i] == 'D' {
			if r == 0 {
				d++
				leftD++
			} else {
				senate[i] = ' '
				r--
			}
		}
	}
	if leftR > 0 && leftD == 0 {
		return "Radiant"
	}
	if leftR == 0 && leftD > 0 {
		return "Dire"
	}
	return ppv(senate, r, d)
}
