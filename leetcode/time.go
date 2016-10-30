package main

import (
	"fmt"
	"time"
	"math"
)

const (
	dayInNanoSecond = 1e9 * 3600 * 24
)

func main() {
	now := time.Now().UnixNano()
	today := now - now%(dayInNanoSecond) - int64(8*3600*math.Pow(10, 9))
	fmt.Println(today)
}
