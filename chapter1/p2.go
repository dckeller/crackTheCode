package main 

import (
	"fmt"
	"strconv"
)

func isPermitation(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	set := make(map[rune]int)

	for _, letter := range word1 {
		set[letter] += 1
		fmt.Printf("Word 1: %v \n", set)
	}

	for _, letter := range word2 {
		set[letter] -= 1
		fmt.Printf("Word 2: %v \n", set)

			if set[letter] < 0 {
				return false
			}
	}

	return true
}

func main() {
	fmt.Println(isPermitation("abc", "acb"))
}