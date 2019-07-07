package main

import "fmt"

func main() {
	flowerbed := []int{1,0,0,0,1}
	n := 1
	fmt.Println(canPlaceFlowers(flowerbed, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	count := 0
	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		return 1 >= n
	}
	if len(flowerbed) > 1 && flowerbed[0] == 0 && flowerbed[1] == 0 {
		count++
		flowerbed[0] = 1
	}
	for i := 1; i < len(flowerbed) - 1; i++ {
		if flowerbed[i] == 1 {
			continue
		}
		if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			count++
		}
		if count >= n {
			return true
		}
	}
	if len(flowerbed) > 1 && flowerbed[len(flowerbed)-2] == 0 && flowerbed[len(flowerbed)-1] == 0 {
		count++
	}
	return count >= n
}
