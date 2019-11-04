package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please input a floating number:")
	var num float32
	fmt.Scan(&num)
	res := int(num)
	fmt.Println(res)
}
