package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Please input a string:")
	var inputStr string
	fmt.Scan(&inputStr)
	inputStr = strings.ToLower(inputStr)

	if inputStr[0] == 'i' && inputStr[len(inputStr)-1] == 'n' && strings.Contains(inputStr, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
