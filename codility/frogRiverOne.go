package main 

import (
	"fmt"
)

func Solution(X int, A []int) int {
	for i, v := range A {
		if v == X {
			return i
		}
	}
	return -1
}

func main() {
	a := []int{1, 3, 1, 4, 2, 3, 5, 4}
	fmt.Println(Solution(5, a))
}