package main

import (
	"fmt"
)

func Solution(N int) int {
	var result 			int
	var tempResult	int
	var calc				bool

	for N > 0 {
		
		// If N ends with a 1
		if N % 2 == 1{
			fmt.Println(N)
			if !calc {
				calc = true
			} else {
				if tempResult > result {
					fmt.Println(tempResult)
					result = tempResult
				}
				tempResult = 0
			}
		} else {
			if calc {
				tempResult += 1
				fmt.Println(tempResult)
			}
		}

		N = N/2
	}
	fmt.Println(result)
	return result
}

func main() {
	Solution(1041)
}