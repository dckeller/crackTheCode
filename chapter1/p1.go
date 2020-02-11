package main 

import (
	"fmt"
)

func isUnique(word string) bool {
	stringlength := len(word)

	for i := 0; i < stringlength; i++ {
		for j := i + 1; j < stringlength; j++ {
				fmt.Println(word[i])
				fmt.Println(word[j])
				fmt.Println("--------------")
			if word[i] == word[j] {
				return false
			}
		}  
	}
	return true
}

func main() {
	fmt.Println(isUnique("word"))
}