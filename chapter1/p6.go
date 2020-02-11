package main

import (
	"fmt"
	"strings"
	"strconv"
)

func stringCompression(letters string) string {
	counter := make(map[rune]int)
	finalString := ""

	for _, char := range letters {
		counter[char] += 1
	}
	
	for keys, values := range counter {
		finalString += strconv.QuoteRune(keys)
		finalString += strconv.Itoa(values)
	}
	finalString = strings.Replace(finalString, "'", "", -1)
	return finalString
}

func main() {
	fmt.Println(stringCompression("aaaabbbbbbbbbccc"))
}