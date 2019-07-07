package main

import "fmt"

func main() {
	fmt.Println(numDupDigitsAtMostN(8766), 4060)
	fmt.Println(numDupDigitsAtMostN(20), 1)
	fmt.Println(numDupDigitsAtMostN(123), 23)
	fmt.Println(numDupDigitsAtMostN(676767), 557337)
}

func numDupDigitsAtMostN(N int) int {
	if N <= 10 {
		return 0
	}
	nums := make([]int, 0)
	t := N
	for t > 0 {
		nums = append(nums, t%10)
		t = t/10
	}
	d := 2
	add := 9
	now := 9
	sum := 9
	for d < len(nums) {
		now *= add
		sum += now
		d++
		add--
	}
	m := make(map[int]bool)
	for i := len(nums) - 1; i >= 0; i-- {
		now = nums[i]
		if i == len(nums) - 1 {
			now--
		} else {
			tmp := 0
			for x := 0; x < now; x++ {
				if !m[x] {
					tmp++
				}
			}
			now = tmp
		}
		add = 10 - (len(nums) - i)
		d = i
		for d > 0 {
			now *= add
			add--
			d--
		}
		sum += now
		if m[nums[i]] {
			break
		}
		m[nums[i]] = true
	}
	if len(m) == len(nums) {
		sum++
	}
	return N - sum
}
