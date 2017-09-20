package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	n := 2
	logs := []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"}
	//logs := []string{"0:start:0", "1:start:2", "2:start:3", "2:end:4", "1:end:5", "0:end:6"}
	fmt.Println(exclusiveTime(n, logs))
}

// queue
func exclusiveTime(n int, logs []string) []int {
	l := list.New()
	ret := make([]int, n)

	m := make(map[int]int)
	mTime := make(map[int]int)
	mPos := make(map[int]int)

	var pos, preVal, nowVal int
	for i := 0; i < len(logs); i++ {
		s := strings.Split(logs[i], ":")
		s = append(s, strconv.Itoa(i))
		if s[1] == "start" {
			l.PushBack(s)
		} else {
			pos, _ = strconv.Atoi(s[0])
			nowVal, _ = strconv.Atoi(s[2])
			now := l.Back()
			val := now.Value.([]string)
			preVal, _ = strconv.Atoi(val[2])

			m[i], _ = strconv.Atoi(val[3])
			mTime[i] = nowVal - preVal + 1
			mPos[i] = pos

			l.Remove(now)
		}
	}
	fmt.Println(m)
	fmt.Println(mTime)
	fmt.Println(mPos)
	for k, v := range m {
		pos = mPos[k]
		fmt.Println("pos", pos, k, v)
		ret[pos] += mTime[k]
		for i := k - 1; i > v; {
			fmt.Println("i:", i)
			if t, ok := mTime[i]; ok {
				fmt.Println("t: ", ret[pos], t)
				ret[pos] -= t
				i = m[i]
			} else {
				i--
			}
		}
	}
	return ret
}

func exclusiveTime1(n int, logs []string) []int {
	if len(logs) <= 1 {
		return []int{}
	}
	m := make(map[string][]int)
	for i := 0; i < n; i++ {
		ii := strconv.Itoa(i)
		m[ii] = []int{}
	}
	s := strings.Split(logs[0], ":")

	pre := s[0]
	preStaus := s[1]
	preVal, _ := strconv.Atoi(s[2])

	m[s[0]] = []int{preVal}

	for i := 1; i < len(logs); i++ {
		fmt.Println("pre:", pre, preStaus, preVal)
		s = strings.Split(logs[i], ":")
		val, _ := strconv.Atoi(s[2])
		fmt.Println("now", s)

		if preStaus != "end" && pre != s[0] {
			m[pre] = append(m[pre], val)
		}

		if s[1] == "end" {
			if pre != s[0] {
				m[s[0]] = append(m[s[0]], preVal+1)
			}
		}

		if s[1] == "end" {
			m[s[0]] = append(m[s[0]], val+1)
		} else {
			m[s[0]] = append(m[s[0]], val)
		}

		pre = s[0]
		preStaus = s[1]
		preVal = val
		fmt.Println(m)
	}

	ret := []int{}
	for i := 0; i < n; i++ {
		ii := strconv.Itoa(i)
		fmt.Println(m[ii])
		value := m[ii]
		sum := 0
		for j := 0; j < len(value); j += 2 {
			sum += value[j+1] - value[j]
		}
		ret = append(ret, sum)
	}
	return ret
}
