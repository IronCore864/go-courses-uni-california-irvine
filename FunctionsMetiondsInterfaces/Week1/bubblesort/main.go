package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(i, j int, slice []int) {
	tmp := slice[i]
	slice[i] = slice[j]
	slice[j] = tmp
}

func BubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if slice[i] > slice[j] {
				swap(i, j, slice)
			}
		}
	}
}

func main() {
	fmt.Println("Please input numbers, maximum 10:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	words := strings.Fields(input)

	slice := make([]int, 0, 10)
	for _, w := range words {
		n, _ := strconv.Atoi(w)
		slice = append(slice, n)
	}

	BubbleSort(slice)

	fmt.Println(slice)
}
