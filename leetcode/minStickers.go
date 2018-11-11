package main

import "fmt"

func main() {
	// stickers := []string{"ac", "db", "ad"}
	// target := "abc"
	//stickers := []string{"fly","me","charge","mind","bottom"}
	// target := "centorder"

	stickers := []string{"control", "heart", "interest", "stream", "sentence", "soil", "wonder", "them", "month", "slip", "table", "miss", "boat", "speak", "figure", "no", "perhaps", "twenty", "throw", "rich", "capital", "save", "method", "store", "meant", "life", "oil", "string", "song", "food", "am", "who", "fat", "if", "put", "path", "come", "grow", "box", "great", "word", "object", "stead", "common", "fresh", "the", "operate", "where", "road", "mean"}
	target := "stoodcrease"
	fmt.Println(minStickers(stickers, target))
}

func minStickers(stickers []string, target string) int {
	fmt.Println(stickers)
	fmt.Println(target)
	tm := make(map[byte]int)
	success := make(map[byte]bool)
	for i := 0; i < len(target); i++ {
		tm[target[i]] += 1
	}
	haveIdx := make(map[byte]map[int]bool)
	source := make([]map[byte]int, 0)
	idx := 0
	for _, s := range stickers {
		sm := make(map[byte]int)
		for i := 0; i < len(s); i++ {
			if tm[s[i]] > sm[s[i]] {
				sm[s[i]]++
				success[s[i]] = true
				if _, ok := haveIdx[s[i]]; !ok {
					haveIdx[s[i]] = make(map[int]bool)
				}
				haveIdx[s[i]][idx] = true
			}
		}
		if len(sm) > 0 {
			source = append(source, sm)
			idx++
		}
	}
	if len(success) < len(tm) {
		return -1
	}
	count := 0
	fmt.Println(haveIdx)
	for k, v := range haveIdx {
		if len(v) > 1 {
			continue
		}
		if tm[k] <= 0 {
			continue
		}
		for k1 := range v {
			fmt.Println("11: ", k1, source[k1])
			for k2, v2 := range source[k1] {
				if tm[k2] > 0 {
					tm[k2] -= v2
					if tm[k2] <= 0 {
						delete(tm, k2)
					}
				}
			}
		}
		count++
	}
	return count + calStickers(source, tm)
}

func calStickers(source []map[byte]int, target map[byte]int) int {
	
}

func calStickers(source []map[byte]int, target map[byte]int) int {
	if len(target) == 0 {
		return 0
	}
	sameCount := 0
	idxs := []int{}
	for i, s := range source {
		count := 0
		for k := range s {
			if target[k] > 0 {
				count += min(target[k], s[k])
			}
		}
		if count > sameCount {
			sameCount = count
			idxs = []int{i}
		} else if count == sameCount {
			idxs = append(idxs, i)
		}
	}
	min := 2 << 31
	fmt.Println(idxs)
	for _, idx := range idxs {
		tmp := make(map[byte]int)
		for k, v := range target {
			tmp[k] = v
		}
		for k, v := range source[idx] {
			if target[k] > 0 {
				target[k] -= v
				if target[k] <= 0 {
					delete(target, k)
				}
			}
		}
		count := calStickers(source, target)
		if count < min {
			min = count
		}
		target = tmp
	}
	return min + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
