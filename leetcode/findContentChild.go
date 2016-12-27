package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1, 2, 3, 4, 5}
	//fmt.Println(findContentChildren(g, s))
	fmt.Println(minMoves(g))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j, nums := 0, 0, 0
	lenG, lenS := len(g), len(s)
	for i < lenG && j < lenS {
		if g[i] <= s[j] {
			i++
			j++
			nums++
		} else {
			j++
		}
	}
	return nums
}

func minMoves(nums []int) int {
	sort.Ints(nums)
	lenN := len(nums)
	sum := 0
	times := 0
	for i := 0; i < lenN; i++ {
		sum += nums[i]
	}
	big := nums[lenN-1] * lenN
	times = (big - sum)/(lenN-1)
	sum += time * (lenN - 1)
	for {
		for sum%lenN != 0 {
			times++
			sum += lenN - 1
		}
		avg := sum/lenN
		last := avg - nums[lenN-1]
		if last < 0 {
			sum += lenN - 1
			times++
			continue
		} else if last == 0 {
			i := lenN - 2
			for i >= 0 {
				if (nums[i] + times) != avg {
					sum += lenN - 1
					times++
					break
				}
				i--
			}
			if i < 0 {
				return times
			} else {
				continue
			}
		} else {
			for i := 0; i < lenN-1; i++ {
				fmt.Println("333 : ", times, avg, last)
				t := times - (avg - nums[i])
				if t < 0 || last < 0{
					sum += lenN - 1
					times++
					break
				}
				last = last - t
			}
			if last == 0 {
				return times
			}
		}
	}

	return 0
}
