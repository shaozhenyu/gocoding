package main

import "fmt"

func main() {
	A := []int{1, 2, 3}
	k := 3
	fmt.Println(largestSumOfAverages(A, k))
}

func largestSumOfAverages(A []int, K int) float64 {
	visited := make([][]float64, len(A)) // [idx][k]int
	sums := make([]int, len(A))
	sum := 0
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]float64, K + 1)
		sum += A[i]
		sums[i] = sum
	}
	return sumOfAverages(A, 0, K, sums, visited)
}

func sumOfAverages(A []int, idx, k int, sum []int, visited [][]float64) float64 {
	if idx == len(A) {
		return 0
	}
	if k == 1 {
		return float64(sum[len(A)-1]-sum[idx]+A[idx]) / float64(len(A)-idx)
	}
	fmt.Println(idx, k)
	if visited[idx][k] > 0 {
		return visited[idx][k]
	}
	var max float64
	for i := idx; i < len(A); i++ {
		now := sum[i] - sum[idx] + A[idx]
		v := float64(now)/float64(i-idx+1) + sumOfAverages(A, i+1, k-1, sum, visited)
		if v > max {
			max = v
		}
	}
	visited[idx][k] = max
	return max
}
