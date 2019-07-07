package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
)

func main() {
	prices := []string{"0.700","2.800","4.900"}
	target := 8
	fmt.Println(minimizeError(prices, target))
}

func minimizeError(prices []string, target int) string {
	sum := 0
	f := []float64{}
	for i := 0; i < len(prices); i++ {
		s := strings.Split(prices[i], ".")
		s1, _ := strconv.Atoi(s[0])
		s2, _ := strconv.ParseFloat(s[1], 10)
		s2 *= 0.001
		sum += s1
		f = append(f, float64(1) - s2)
	}
	if sum > target || sum + len(prices) < target{
		return "-1"
	}
	sum += len(prices)
	sort.Float64s(f)
	t := sum - target
	var ret float64
	for i := 0; i < len(prices) - t; i++ {
		ret += f[i]
	}
	fmt.Println(f)
	for i := len(prices) - t; i < len(prices); i++ {
		ret += (float64(1) - f[i])
	}
	return fmt.Sprintf("%.3f", ret)
}
