package main 

import (
	"fmt"
	"strings"
	"strconv"
)

// func HeapPermutation(a []string, size int) {
// 	fmt.Println(size)

// 	if size == 1 {
// 		fmt.Println(a)
// 	}

// 	for i := 0; i < size; i++ {
// 		HeapPermutation(a, size-1)

// 		if size%2 == 1 {
// 			a[0], a[size-1] = a[size-1], a[0]
// 		} else {
// 			a[i], a[size-1] = a[size-1], a[i]
// 		}
// 		fmt.Println(size)
// 		fmt.Println("----------")
// 	}

// }

func UniquePerm(permList []string, check string) bool {
	for _, val := range permList {
		if val == check {
			return false
		}
	}
	return true
}

func CreatePerms(input []string, size int) int {
	var uniqueList []string
	var currentString string
	

	for i := 0; i < size; i++ {
		currentString += input[i]

		// Check if unique
		if UniquePerm(uniqueList, currentString) == true {
			uniqueList = append(uniqueList, currentString)
			input[0], input[size-1] = input[size-1], input[0]
		} else {
			// input[i], input[size-1] = input[size-1], input[i]
			fmt.Println(input)
		}
		
		fmt.Println(currentString)
		currentString = ""
	}

	return len(uniqueList)
}

func main() {
	str := strconv.Itoa(123)
  a := strings.Split(str, "")
  // HeapPermutation(str, len(a))
  CreatePerms(a, len(a))
}