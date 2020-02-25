package main 

import (
	"fmt"
)

func Solution(A []int) int {
	compareArray := make([]int, 0)
	sumA := 0
	comA := 0

	for i := range A {
		compareArray = append(compareArray, i + 1)
	}

	// Find sum of given array 
	for _, vals := range A {
		sumA += vals
	}

	// Find sum of given array 
	for _, vals := range compareArray {
		comA += vals
	}


	if len(A) == len(compareArray) {
		if sumA > comA {
			return sumA - comA
		} else {
			return comA - sumA
		}
	} 

	// if A[i] < 0 {
	// 	fmt.Println(1)
	// 	return 1
	// } else {
	// 	fmt.Println(len(A) + 1)
	// 	return len(A) + 1
	// }
return -1
}

func main() {
	a := []int{1, 3, 6, 4, 1, 2}
	fmt.Println(Solution(a))
}