package main 

import (
	"fmt"
)

func Solution(x, y, d int) int {
	counter := 0

	for i := x; i <= y; i+=d {
		counter += 1
	}
	fmt.Println(counter)
	return counter
}

func main() {
	Solution(10, 85, 30)
}