package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	max := height[0]
	idx := 0
	for i := 1; i < len(height); i++ {
		if max < height[i] {
			max = height[i]
			idx = i
		}
	}
	return trap1(height[:idx]) + trap2(height[idx+1:])
}

func trap1(h []int) int {
	if len(h) == 0 {
		return 0
	}
	max := h[0]
	idx := 0
	for i := 1; i < len(h); i++ {
		if max < h[i] {
			max = h[i]
			idx = i
		}
	}
	ret := 0
	for i := idx + 1; i < len(h); i++ {
		ret += max - h[i]
	}
	return ret + trap1(h[:idx])
}

func trap2(h []int) int {
	if len(h) == 0 {
		return 0
	}
	max := h[0]
	idx := 0
	for i := 1; i < len(h); i++ {
		if max < h[i] {
			max = h[i]
			idx = i
		}
	}
	ret := 0
	for i := idx - 1; i >= 0; i-- {
		ret += max - h[i]
	}
	return ret + trap2(h[idx+1:])
}
