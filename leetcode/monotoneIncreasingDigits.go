package main

import "fmt"

func main() {
	fmt.Println(monotoneIncreasingDigits(123456))
	fmt.Println(monotoneIncreasingDigits(1023456))
	fmt.Println(monotoneIncreasingDigits(7123456))
	fmt.Println(monotoneIncreasingDigits(668841))
}

func monotoneIncreasingDigits(N int) int {
	arr := []int{}
	for N > 0 {
		d := N%10
		arr = append(arr, d)
		N = N/10
	}
	lenN := len(arr)
	for i := 0; i < lenN/2; i++ {
		arr[i], arr[lenN-1-i] = arr[lenN-1-i], arr[i]
	}
	for i := 1; i < lenN; i++ {
		if arr[i-1] <= arr[i] {
			continue
		} else {
			j := i - 2
			for ; j >= 0; j-- {
				fmt.Println(arr[j], arr[i-1])
				fmt.Println(j, i-1)
				if arr[j] != arr[i-1] {
					break
				}
			}
			fmt.Println(i, j)
			j += 1
			i = j + 1
			arr[j] -= 1
			for i < lenN {
				arr[i] = 9
				i++
			}
			break
		}
	}
	ret := 0
	for i := 0; i < lenN; i++ {
		ret = ret * 10 + arr[i]
	}
	fmt.Println(arr)
	return ret
}
