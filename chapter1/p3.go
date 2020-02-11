package main 

import (
	"fmt"
	"strings"
)

func insertString(sentence string) string {
	splitString := strings.Split(sentence, " ")

	joined := strings.Join(splitString, "%203")

	return joined
}

func main() {
	fmt.Println(insertString("Mr 3ohn Smith"))
}