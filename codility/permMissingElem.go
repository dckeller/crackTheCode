package main 

import (
	"fmt"
)

func Solution(A []int) int {
	compareArray := make([]int, len(A) + 1)
	length := len(compareArray)
	sum := 0
	compSum := 0

	// Find sum of given array 
	for _, vals := range A {
		sum += vals
	}

	// Create compare array
	for i := 0; i < length; i++ {
		compareArray = append(compareArray, i + 1)
	}

	// Find some of compare array
	for _, vals := range compareArray {
		compSum += vals
	}

	return (compSum - sum)
}

func main() {
	a := []int{2, 3, 1, 5}
	Solution(a)
}