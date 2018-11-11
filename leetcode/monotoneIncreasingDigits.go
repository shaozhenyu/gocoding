package main

import "fmt"

func main() {
	fmt.Println(monotoneIncreasingDigits(123456), monotoneIncreasingDigits1(123456))
	fmt.Println(monotoneIncreasingDigits(1023456), monotoneIncreasingDigits1(1023456))
	fmt.Println(monotoneIncreasingDigits(7123456), monotoneIncreasingDigits1(7123456))
	fmt.Println(monotoneIncreasingDigits(668841), monotoneIncreasingDigits1(668841))
}

func monotoneIncreasingDigits(N int) int {
	if N < 10 {
		return N
	}
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
	doNext(0, arr)
	ret := 0
	for i := 0; i < lenN; i++ {
		ret = ret * 10 + arr[i]
	}
	return ret
}

func doNext(idx int, n []int) {
	if len(n) == idx + 1 {
		return
	}
	if n[idx] <= n[idx+1] {
		doNext(idx+1, n)
	}
	if n[idx] > n[idx+1] {
		n[idx]--
		for i := idx+1; i < len(n); i++ {
			n[i] = 9
		}
	}
}

func monotoneIncreasingDigits1(N int) int {
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
				if arr[j] != arr[i-1] {
					break
				}
			}
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
