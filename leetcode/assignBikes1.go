package main

import "fmt"

func main() {

}

func assignBikes(workers [][]int, bikes [][]int) int {
    min := 2 << 31
    ab(0, 0, &min, workers, bikes, make([]bool, len(bikes)))
    return min
}

func ab(wIdx, total int, min *int, workers, bikes [][]int, used []bool) {
    if wIdx == len(workers) {
        if total < *min {
            *min = total
            return
        }
    }
    for i := 0; i < len(bikes); i++ {
        if used[i] {
            continue
        }
        used[i] = true
        ab(wIdx+1, total + Manhattan(workers[wIdx], bikes[i]), min, workers, bikes, used)
        used[i] = false
    }
}

func Manhattan(a, b []int) int {
    return abs(a[0] - b[0]) + abs(a[1] - b[1])
}

func abs(x int) int {
    if x < 0 {
        x *= -1
    }
    return x
}
