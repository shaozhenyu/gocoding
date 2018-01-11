package main

import "fmt"

func main() {
	nums := []int{9, 4, 1, 3, 6, 7, 5, 8, 0, 2}
	quicksort(nums, 0, 9)
	fmt.Println(nums)
}

func quicksort(nums []int, low, high int) {
	if low < high {
		i := low
		j := high
		temp := nums[i]

		for i < j {
			for nums[j] >= temp && i < j {
				j--
			}
			nums[i] = nums[j]
			for nums[i] <= temp && i < j {
				i++
			}
			nums[j] = nums[i]

			nums[i] = temp
			quicksort(nums, low, i-1)
			quicksort(nums, j+1, high)
		}
	}
}
