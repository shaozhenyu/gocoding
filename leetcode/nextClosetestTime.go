package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(nextClosestTime("23:49"))

}

func nextClosestTime(time string) string {

	t := [4]int{}

	t[0], _ = strconv.Atoi(string(time[0]))
	t[1], _ = strconv.Atoi(string(time[1]))
	t[2], _ = strconv.Atoi(string(time[3]))
	t[3], _ = strconv.Atoi(string(time[4]))

	oldHour := t[0]*10 + t[1]
	oldMin := t[2]*10 + t[3]

	closest := 24 * 60
	closestTime := ""

	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			for c := 0; c < 4; c++ {
				for d := 0; d < 4; d++ {
					din := 0
					hour := t[a]*10 + t[b]
					min := t[c]*10 + t[d]

					if hour == oldHour && min == oldMin {
						continue
					}

					if hour < 0 || hour >= 24 {
						continue
					}
					if min < 0 || min >= 60 {
						continue
					}

					if hour > oldHour {
						din = (hour - oldHour) * 60
						din += min - oldMin
					}

					if hour == oldHour {
						if min > oldMin {
							din = min - oldMin
						} else {
							din = 24*60 + min - oldMin
						}
					}

					if hour < oldHour {
						din = (hour + 24 - oldHour) * 60
						din += min - oldMin
					}

					if din < closest {
						var strH, strMin string
						strH = strconv.Itoa(hour)
						if hour < 10 {
							strH = "0" + strH
						}
						strMin = strconv.Itoa(min)
						if min < 10 {
							strMin = "0" + strMin
						}

						closestTime = strH + ":" + strMin
						closest = din
					}
				}
			}
		}
	}

	return closestTime
}
