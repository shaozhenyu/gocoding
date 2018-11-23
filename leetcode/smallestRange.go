package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		[]int{1, 2, 3},
	}
	fmt.Println(smallestRange(nums))
}

// 2018 11 15
func smallestRange(nums [][]int) []int {
	bucket := make([]map[int]struct{}, 200001)
	start := 200000
	end := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			v := 100000 + nums[i][j]
			if v < start {
				start = v
			}
			if v > end {
				end = v
			}
			if bucket[v] == nil {
				bucket[v] = make(map[int]struct{})
			}
			bucket[v][i] = struct{}{}
		}
	}
	size := len(nums)
	ret := []int{start, end}
	d := end - start
	min := start
	flag := false
	fmt.Println(bucket[start:end+1])
	exists := make(map[int]int)
	for i := start; i <= end; i++ {
		if bucket[i] == nil {
			continue
		}
		for k := range bucket[i] {
			if _, ok := exists[k]; !ok {
				flag = true
			}
			if i > exists[k] && exists[k] == min {
				flag = true
			}
			exists[k] = i
		}
		if len(exists) == size && flag {
			flag = false
			a := i
			for _, v := range exists {
				if v < a {
					a = v
				}
			}
			if a >= min {
				min = a
				if i - min < d {
					ret = []int{min, i}
					d = i - min
				}
			}
		}
	}
	return []int{ret[0]-100000, ret[1]-100000}
}



// old
type NeedSort [][]int

func (n NeedSort) Len() int {
	return len(n)
}

func (n NeedSort) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n NeedSort) Less(i, j int) bool {
	if n[i][0] != n[j][0] {
		return n[i][0] < n[j][0]
	} else {
		return n[i][1] < n[j][1]
	}
	return false
}

func smallestRange1(nums [][]int) []int {
	//merge k
	all := [][]int{}
	k := len(nums)

	if k == 1 {
		return []int{nums[0][0], nums[0][0]}
	}

	for i := 0; i < k; i++ {
		for j := 0; j < len(nums[i]); j++ {
			this := []int{nums[i][j], i}
			all = append(all, this)
		}
	}
	sort.Sort(NeedSort(all))

	m := make(map[int]int)

	shortest := 100000 * 2
	short := make([]int, 2)
	start := all[0]
	end := all[len(all)-1]
	m[all[0][1]] = start[0]
	for i := 1; i < len(all); i++ {
		end = all[i]
		if _, ok := m[all[i][1]]; !ok {
			m[all[i][1]] = all[i][0]
			if len(m) >= k {
				lenNow := end[0] - start[0]
				if lenNow < shortest {
					shortest = lenNow
					short[0], short[1] = start[0], end[0]
					if shortest == 0 {
						return short
					}
				}

				delete(m, start[1])
				smallest := all[i][0]
				key := -1
				for k, v := range m {
					if v <= smallest {
						smallest = v
						key = k
					}
				}
				if key != -1 {
					start[0], start[1] = m[key], key
				}
			}
			if len(m) >= k && start[1] == all[i][0] {
				smallest := all[i][0]
				key := 0
				for k, v := range m {
					if v < smallest {
						smallest = v
						key = k
					}
				}
				start[0], start[1] = m[key], key
				delete(m, key)
			}
		} else {
			m[all[i][1]] = all[i][0]
			if len(m) == 1 {
				start[0], start[1] = all[i][0], all[i][1]
			} else if all[i][1] == start[1] {
				smallest := all[i][0]
				key := -1
				for k, v := range m {
					if v <= smallest {
						smallest = v
						key = k
					}
				}
				if key != -1 {
					start[0], start[1] = m[key], key
				}
			}
		}
	}

	return short
}
