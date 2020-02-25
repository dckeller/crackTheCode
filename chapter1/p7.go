package main 

import (
	"fmt"
	"math/rand"
)

func main() {
	var array = [2][2]int{}
	dim := len(array)

	for i := 0; i < dim; i ++ {
		for j := 0; j < dim; j++ {
			array[i][j] = rand.Intn(10)
			fmt.Printf("a[%d][%d] = %d\n", i,j, array[i][j] )
		}
	}
}