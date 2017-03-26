package main

import (
	"strings"
	"strconv"
	"sort"
	"fmt"
)

func main() {
	ret := findMinDifference([]string{"00:00", "22:30", "00:00"})
	println(ret)
}

func findMinDifference(timePoints []string) int {
	lenT := len(timePoints)
	if lenT <= 1 {
		return 0
	}
	arr := make([]int, lenT)
	for i := 0; i < lenT; i++ {
		times := strings.Split(timePoints[i], ":")
		hour, _ := strconv.Atoi(times[0])
		minute, _ := strconv.Atoi(times[1])
		arr[i] = hour * 60 + minute
	}
	
	sort.Ints(arr)
	fmt.Println(arr)
	min := 60 * 24
	for i := 1; i < lenT; i++ {
		tmp := arr[i] - arr[i-1]
		if min > tmp {
			min = tmp
		}
	}
	tmp := 60 * 24 - arr[lenT-1] + arr[0]
	if min > tmp  {
		min = tmp
	}
	return min
}
