package main

import (
	"fmt"
) 

func Solution(A []int) int {
	counter := make(map[int]int)

	for _, val := range A {
		counter[val] += 1
	}

	for key, values := range counter {
		if values == 1 {
			fmt.Println(key)
			return key
		}
	}
	return -1
}

func main() {
	a := []int{9, 3, 9, 3, 9, 7, 9}
	Solution(a)
}