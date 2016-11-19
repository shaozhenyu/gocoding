package main

func main() {
	a := []int{1, 3}
	b := []int{2, 4}
	println(findMedianSortedArrays(a, b))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	i, j, k := 0, 0, 0

	nums3 := make([]int, len1+len2)

	for i < len1 && j < len2 {
		if nums1[i] <= nums2[j] {
			nums3[k] = nums1[i]
			i++
		} else {
			nums3[k] = nums2[j]
			j++
		}
		k++
	}

	for i < len1 {
		nums3[k] = nums1[i]
		k++
		i++
	}
	for j < len2 {
		nums3[k] = nums2[j]
		k++
		j++
	}
	println(k)
	if k%2 == 0 {
		return (float64(nums3[k/2]) + float64(nums3[k/2-1]))/2
	} else {
		return float64(nums3[k/2])
	}
	return 0
}
