package main 

import (
	"fmt"
)

func Solution(N int, A []int) []int {
	counterArray := make([]int, N)
	count := 0

	for _, v := range A {
		if v <= N {
			counterArray[v - 1] += 1
			for i := 0; i < len(counterArray); i++ {
				if counterArray[i] > count {
					count = counterArray[i]
				}
			}
		} else {
			for j := 0; j < len(counterArray); j++ {
				counterArray[j] = count
			} 
		}
	}
	return counterArray
}

func main() {
	a := []int{3, 4, 4, 6, 1, 4, 4}
	Solution(5, a)
}