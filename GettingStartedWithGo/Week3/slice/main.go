package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 0, 3)
	fmt.Println("Please input numbers, X to stop.")
	var input string
	fmt.Scan(&input)

	for input != "X" {
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			fmt.Scan(&input)
			continue
		}
		sli = append(sli, num)
		fmt.Println("Please input numbers, X to stop.")
		fmt.Scan(&input)
	}
	sort.Ints(sli)
	fmt.Println("After sort")
	fmt.Println(sli)
}
