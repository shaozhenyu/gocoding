package main

import "fmt"

func main() {
	N := 5
	K := 12
	fmt.Println(kthGrammar(N, K))
}

//1 0
//2 01
//3 0110
//4 0110 1001
//5 0110 1001 1001 0110
//6 0110 1001 1001 0110 1001 0110 0110 1001
func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	total := 1
	line := N
	for N > 1 {
		total *= 2
		N--
	}
	flag := true
	for line > 0 && K > 1 {
		fmt.Println(line, total, K)
		for K <= total {
			total /= 2
			line--
		}
		// 对称
		if line%2 == 0 {
		} else { // 相反
			flag = !flag
		}
		K = total - (K - total - 1)
	}
	if flag {
		return 0
	}
	return 1
}
