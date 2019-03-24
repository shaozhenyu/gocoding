package main

import "fmt"

func main() {
	poured := 5
	query_row := 2
	query_glass := 1
	fmt.Println(champagneTower(poured, query_row, query_glass))
}

func champagneTower(poured int, query_row int, query_glass int) float64 {
	all := make([][]float64, query_row+1)
	for i := 0; i < query_row+1; i++ {
		all[i] = make([]float64, query_row+1)
	}
	all[0][0] = float64(poured)
	for i := 0; i < query_row; i++ {
		for j := 0; j <= i; j++ {
			if all[i][j] <= 1.0 {
				continue
			}
			mid := (all[i][j] - 1)/2.0
			all[i+1][j] += mid
			all[i+1][j+1] += mid
		}
	}
	if all[query_row][query_glass] > 1 {
		all[query_row][query_glass] = 1
	}
	return all[query_row][query_glass]
}
