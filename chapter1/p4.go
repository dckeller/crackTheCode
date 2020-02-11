package main 

import (
	"fmt"
)

func isPalidrome(word string) bool {

	// to count the occurance of chars excluding spaces
	counter := make(map[rune]int)
	odd := false


	for _, char := range word {
		if char == ' ' {
			continue
		}

		// Add chars to the map and count each occurance of letter
		counter[char] += 1
	}

	// iterate through map
	for _, nums := range counter {

		fmt.Println(nums)
		if nums % 2 != 0 {
			if odd {
				return false
			}

			odd = true
		}
	}
	fmt.Println(counter)
	return true
}

func main() {
	fmt.Println(isPalidrome("boiob"))
}