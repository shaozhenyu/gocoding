package main

import "fmt"

func main() {
	words := []string{"oath","pea","eat","rain"}
	board := [][]byte{
		[]byte{'o','a','a','n'},
		[]byte{'e','t','a','e'},
		[]byte{'i','h','k','r'},
		[]byte{'i','f','l','v'},
	}

	board = [][]byte{[]byte{'a','b'},[]byte{'c','d'}}
	words = []string{"ab","cb","ad","bd","ac","ca","da","bc","db","adcb","dabc","abb","acb"}

	board = [][]byte{[]byte{'a','b'},[]byte{'a','a'}}
	// words = []string{"aba","baa","bab","aaab","aaa","aaaa","aaba"}
	words = []string{"aaba"}

	fmt.Println(findWords(board, words))
}


var next [][]int = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return []string{}
	}
	wordM := make(map[byte][]string)
	for i := 0; i < len(words); i++ {
		if _, ok := wordM[words[i][0]]; !ok {
			wordM[words[i][0]] = make([]string, 0)
		}
		wordM[words[i][0]] = append(wordM[words[i][0]], words[i])
	}
	visited := make([][]bool, len(board))
	fail := make([][]map[string]bool, len(board))
	for i := 0; i < len(board); i++ {
		fail[i] = make([]map[string]bool, len(board[i]))
		visited[i] = make([]bool, len(board[i]))
		for j := 0; j < len(board[i]); j++ {
			fail[i][j] = make(map[string]bool)
		}
	}
	success := make(map[string]bool)
	fw(board, wordM, fail, success, visited)
	ret := []string{}
	for s := range success {
		ret = append(ret, s)
	}
	return ret
}

func fw(board [][]byte, words map[byte][]string, fail [][]map[string]bool, success map[string]bool, visited [][]bool) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if strs, ok := words[board[i][j]]; ok  {
				for _, str := range strs {
					if !fail[i][j][str] && !success[str] {
					// if !success[str] {
						visited[i][j] = true
						ok := backtrack(i, j, board, str, []byte(str)[1:], fail, visited)
						visited[i][j] = false
						if ok {
							success[str] = true
						}
					}
				}
			}
		}
	}
}

func backtrack(x, y int, board [][]byte, ori string, word []byte, fail [][]map[string]bool, visited [][]bool) bool {
	fmt.Println(x, y, ori, string(word), fail, visited)
	if len(word) == 0 {
		return true
	}
	for _, n := range next {
		i := x + n[0]
		j := y + n[1]
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			continue
		}
		if visited[i][j] {
			continue
		}
		if board[i][j] != word[0] {
			continue
		}
		if fail[i][j][string(word)] {
			continue
		}
		visited[i][j] = true
		ok := backtrack(i, j, board, ori, word[1:], fail, visited)
		visited[i][j] = false
		if ok {
			return ok
		}
		fail[i][j][string(word)] = true
	}
	return false
}
