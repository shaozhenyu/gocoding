package main

import "fmt"

func main() {
	a := make([]int, 2, 1000)
	a[0] = 2
	b := []int{1}
	merge(a, 1, b, 1)
	fmt.Println(a)
}


func merge(nums1 []int, m int, nums2 []int, n int) {

	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}


	k := len(nums1) - 1

	i := m - 1
	j := n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}

func merge1(nums1 []int, m int, nums2 []int, n int) {

	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] == 0 {
			break
		}
		if nums1[i] > nums2[j] {
			rear := make([]int, len(nums1[i:]))
			copy(rear, nums1[i:])
			nums1 = append(nums1[:i], nums2[j])
			nums1 = append(nums1, rear...)
			j++
			m++
		} else {
			i++
		}
	}

	if j < n {
		nums1 = append(nums1[:i], nums2[j:]...)
	}
}
