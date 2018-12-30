package main

import "fmt"

func main() {
	words := []string{"cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"}
	fmt.Println(findAllConcatenatedWordsInADict(words))
}

func findAllConcatenatedWordsInADict(words []string) []string {
    m := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		m[words[i] = true
	}
}
