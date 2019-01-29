package main

import "fmt"

func main() {
	asteroids := []int{-2, -2, 1, -1}
	fmt.Println(asteroidCollision(asteroids))
}

func asteroidCollision(asteroids []int) []int {
	if len(asteroids) <= 1 {
		return asteroids
	}
	left := 0
	right := 1
	for right < len(asteroids) {
		if left >= 0 && asteroids[left] > 0 && asteroids[right] < 0 {
			r := abs(asteroids[right])
			if r > asteroids[left] {
				left--
			} else if r == asteroids[left] {
				left--
				right++
			} else {
				right++
			}
		} else {
			left++
			asteroids[left] = asteroids[right]
			right++
		}
	}
	return asteroids[:left+1]
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}
