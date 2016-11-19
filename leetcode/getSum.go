package main

import "strconv"

func main() {
	println(getSum1(-11, 8))
}


func getSum1(a int, b int) int {
	if b != 0 {
		c := a ^ b
		d := (a & b ) << 1
		return getSum1(c, d)
	} else {
		return a
	}
}
func getSum(a int, b int) int {
	if a >= 0 && b >= 0 {
		a1 := strconv.FormatInt(int64(a), 2)
		b1 := strconv.FormatInt(int64(b), 2)
		add := 0
		lenA := len(a1)
		lenB := len(b1)
		i := lenA - 1
		j := lenB - 1
		str := ""
		for i >= 0 && j >= 0 {
			sum := int(a1[i]) + int(b1[j]) + add - 48 * 2
			if sum <= 1 {
				str = strconv.FormatInt(int64(sum), 2) + str
				add = 0
			} else if sum == 2 {
				str = "0" + str
				add = 1
			} else if sum == 3 {
				str = "1" + str
				add = 1
			}
			i--
			j--
		}

		if i >= 0 {
			sum := int(a1[i]) + add - 48
			if sum <= 1 {
				str = strconv.FormatInt(int64(sum), 2) + str
				add = 0
			} else {
				str = "0" + str
				add = 1
			}
		}
		if j >= 0 {
			sum := int(a1[j]) + add - 48
			if sum <= 1 {
				str = strconv.FormatInt(int64(sum), 2) + str
				add = 0
			} else {
				str = "0" + str
				add = 1
			}
		}
		if add == 1 {
			str = "1" + str
		}
		ret, err := strconv.ParseInt(str, 0, 64)
		if err != nil {
			println(err)
			return 0
		}
		println(ret)
		return int(ret)
	}

	return 0
}
