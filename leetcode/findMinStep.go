package main

import "fmt"

func main() {
	//board := "RBYYBBRRB"
	//hand := "YRBGB"
	//board := "WWRRBBW"
	//hand := "RB"
	//board := "WWRRBBWW"
	//hand := "WRBRW"
	board := "GGRRGRRYY"
	hand := "RYRBR"
	fmt.Println(findMinStep(board, hand))
}

// WWWRW 这种情况不可能发生，因为祖玛发现有3+个连着的会直接消除
func findMinStep(board string, hand string) int {
	board = removeDup(board)
	mHand := make(map[byte]int)
	for i := 0; i < len(hand); i++ {
		mHand[hand[i]]++
	}
	fmt.Println("aa: ", board)
	max := 6
	backTrace(board, mHand, 0, &max)
	if max == 6 {
		return -1
	}
	return max
}

func backTrace(board string, hand map[byte]int, count int, max *int) {
	fmt.Println(board)
	if len(board) == 0 {
		fmt.Println("kkk: ", count)
		*max = min(count, *max)
		return
	}
	for i := 0; i < len(board); {
		j := i + 1
		for j < len(board) && board[i] == board[j] {
			j++
		}
		//if j == len(board) && j-i >= 3 {
		//	fmt.Println("sss: ", count, board)
		//	backTrace(board[:i], hand, count, max)
		//	return
		//}
		if j-i < 3 {
			need := 3 - (j - i)
			if hand[board[i]] >= need {
				hand[board[i]] -= need
				backTrace(removeDup(board[:i]+board[j:]), hand, count+need, max)
				hand[board[i]] += need
			}
		}
		i = j
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// board := "GGRRGRRYY"
func removeDup(board string) string {
	for i, j := 0, 0; i < len(board) && j <= len(board); j++ {
		if j == len(board) {
			if j-i >= 3 {
				return board[:i]
			}
			return board
		}
		if board[i] == board[j] {
			continue
		} else if j-i >= 3 {
			return removeDup(board[:i] + board[j:])
		} else {
			i = j
		}
	}
	return board
}
