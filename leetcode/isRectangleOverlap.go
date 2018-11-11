package main

import "fmt"

func main() {
	rec1 := []int{0, 0, 1, 1}
	rec2 := []int{1, 0, 2, 1}
	fmt.Println(isRectangleOverlap(rec1, rec2))
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return calOverlapping(rec1[0], rec1[2], rec2[0], rec2[2]) > 0 && calOverlapping(rec1[1], rec1[3], rec2[1], rec2[3]) > 0
}

func calOverlapping(A, C, E, G int) int {
    if C < A {
        A, C = C, A
    }
    if G < E {
        G, E = E, G
    }
    if G <= A || E >= C {
        return 0
    }

    if E <= A && G >= C {
        return C - A
    }
    if E >= A && G <= C {
        return G - E
    }
    if E < A {
        return G - A
    } else {
        return C - E
    }
    return 0
}

func abs(a int) int {
    if a < 0 {
        return a * -1
    }
    return a
}
