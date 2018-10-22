package main

import "fmt"

func main() {
	strs := []string{"", "eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	ret := [][]string{}
	m := make(map[int]map[int]map[byte]int)

	for _, str := range strs {
		mStr := make(map[byte]int)
		for i := 0; i < len(str); i++ {
			if _, ok := mStr[str[i]]; ok {
				mStr[str[i]]++
			} else {
				mStr[str[i]] = 1
			}
		}

		if m1, ok := m[len(str)]; ok {
			for index, m2 := range m1 {
				if compare(m2, mStr) {
					ret[index] = append(ret[index], str)
					goto Continue
				}
			}
		}

		ret = append(ret, []string{str})
		if _, ok := m[len(str)]; ok {
			m[len(str)][len(ret)-1] = mStr
		} else {
			newM := make(map[int]map[byte]int)
			newM[len(ret)-1] = mStr
			m[len(str)] = newM
		}
	Continue:
	}

	return ret
}

func compare(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
