package main

import "fmt"

func main() {
	data := []int{197, 130, 1}
	data = []int{235, 140, 4}
	data = []int{240,162,138,147,145}
	// data = []int{240,162,138,147}
	// data = []int{228,189,160,229,165,189,13,10}
	fmt.Println(validUtf8(data))
}

//   Char. number range  |        UTF-8 octet sequence
//      (hexadecimal)    |              (binary)
//   --------------------+---------------------------------------------
//   0000 0000-0000 007F | 0xxxxxxx
//   0000 0080-0000 07FF | 110xxxxx 10xxxxxx
//   0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
//   0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
func validUtf8(data []int) bool {
	if len(data) == 0 {
		return false
	}
	count := 0
	for i := 0; i < len(data); i++ {
		if count == 0 {
			count = getCount(data[i])
		} else {
			if data[i] & (1 << 7) == 0 {
				fmt.Println(data[i] & (1 << 8))
				return false
			}
			count--
		}

		if count == -1 {
			return false
		}
		fmt.Println(fmt.Sprintf("%08b", data[i]), count)
	}
	return true
}

func getCount(n int) int {
	v := fmt.Sprintf("%08b", n)
	fmt.Println("ss: ", v)
	if v[0] == '0' {
		return 0
	}
	if v[:2] == "10" {
		return -1
	}
	if v[:3] == "110" {
		return 1
	}
	if v[:4] == "1110" {
		return 2
	}
	if v[:5] == "11110" {
		return 3
	}
	return -1
}
