package main 

import (
	"fmt"
)

func Solution(A []int, K int) []int {

	for i := 0; i < K; i++ {

		// Grab last int
		grabLastint := A[len(A)-1]

		// Remove last int
		A = A[:len(A)-1]

		// Prepend to front of slice 
		A = append([]int{grabLastint}, A...)
	}
	
	return A
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	Solution(a, 3)
}