package main

import "fmt"

//add pre:  []
//add:  44 53 [{44 53}]
//add pre:  [{44 53}]
//add:  69 89 [{44 53} {69 89}]
//remove pre:  [{44 53} {69 89}]
//remove:  69 91 [{44 86}]
//add pre:  [{44 86}]
//
//add:  87 94 [{44 86} {87 94}]
//remove pre:  [{44 86} {87 94}]
//remove:  34 52 [{52 86} {87 94}]
//add pre:  [{52 86} {87 94}]
//add:  1 59 [{1 86} {87 94}]
//remove pre:  [{1 86} {87 94}]
//remove:  87 96 [{1 62}]
//remove pre:  [{1 62}]
//remove:  62 83 [{1 34}]
//remove pre:  [{1 34}]
//remove:  34 61 [{1 9}]
//
func main() {
	r := Constructor()
	r.AddRange(44, 53)
	r.AddRange(69, 89)
	r.RemoveRange(69, 91)
	fmt.Println(r.itl)
	r.AddRange(87, 94)
	r.RemoveRange(34, 52)
	r.AddRange(1, 59)
	r.RemoveRange(87, 96)
	r.RemoveRange(62, 83)
	r.RemoveRange(34, 61)
	fmt.Println(r.itl)
}

type RangeModule struct {
	itl []interval
}

type interval struct {
	start int
	end   int
}

func Constructor() RangeModule {
	return RangeModule{
		itl: make([]interval, 0),
	}
}

func (this *RangeModule) AddRange(left int, right int) {
	itl := this.itl
	defer func() {
		this.itl = itl
	}()

	for i := 0; i < len(itl); i++ {
		if right < itl[i].start {
			t := append([]interval{}, itl[i:]...)
			itl = append(itl[:i], interval{left, right})
			itl = append(itl, t...)
			return
		}
		if right <= itl[i].end {
			itl[i].start = min(left, itl[i].start)
			return
		}
		if left <= itl[i].end && right > itl[i].end {
			left = min(left, itl[i].start)
			itl = append(itl[:i], itl[i+1:]...)
			i--
		}
	}
	itl = append(itl, interval{left, right})
	return
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	for i := 0; i < len(this.itl); i++ {
		if left >= this.itl[i].end {
			continue
		}
		if left >= this.itl[i].start {
			if right <= this.itl[i].end {
				return true
			} else {
				left = this.itl[i].end
				continue
			}
		}
		if right <= this.itl[i].end {
			break
		}
	}
	return false
}

func (this *RangeModule) RemoveRange(left int, right int) {
	itl := this.itl
	defer func() {
		this.itl = itl
	}()

	for i := 0; i < len(itl); i++ {
		if right <= itl[i].start {
			break
		}
		if left >= itl[i].end {
			continue
		}
		if left <= itl[i].start && right <= itl[i].end {
			itl[i].start = right
			if right == itl[i].end {
				itl = append(itl[:i], itl[i+1:]...)
			}
			return
		}
		if left > itl[i].start && right <= itl[i].end {
			end := itl[i].end
			itl[i].end = left
			if right == end {
				return
			}
			t := append([]interval{}, itl[i+1:]...)
			itl = append(itl[:i+1], interval{right, end})
			itl = append(itl, t...)
			return
		}
		if right > itl[i].end {
			if left <= itl[i].start {
				left = itl[i].end
				itl = append(itl[:i], itl[i+1:]...)
				i--
			} else {
				itl[i].end, left = left, itl[i].end
			}
		}
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
