package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := Constructor(1, 1, 1)
	fmt.Println(s.RandPoint())
}

type Solution struct {
	rd *rand.Rand
	radius float64
	x, y float64
}


func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		rd: rand.New(rand.NewSource(time.Now().UnixNano())),
		radius: radius,
		x: x_center,
		y: y_center,
	}
}


func (this *Solution) RandPoint() []float64 {
	radius := this.radius
	for {
		x := 2 * this.rd.Float64() * radius - radius
		y := 2 * this.rd.Float64() * radius - radius
		if x * x + y * y <= radius * radius {
			return []float64{x + this.x, y + this.y}
		}
	}
	return []float64{}
}

