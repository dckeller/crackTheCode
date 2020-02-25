package main 

import (
	"fmt"
)

func Solution(A []int) int {
	length := len(A) - 1
	sumLeft := 0
	sumright := 0
	tempResult := 0
	result := 1 << 32

for i := 0; i < length; i++ {
		leftSide := A[:i + 1]
		rightSide := A[i + 1:]

	for _, val := range leftSide {
		sumLeft += val
	}

	for _, val := range rightSide {
		sumright += val
	}

	if sumLeft > sumright {
		tempResult = (sumLeft - sumright)
	} else {
		tempResult = (sumright - sumLeft)
	}

	if tempResult < result {
		result = tempResult
	}
	sumright = 0
	sumLeft = 0
}

	fmt.Println(result)
	return result
}

func main() {
	a := []int{3, 1, 2, 4, 3}
	Solution(a)
}