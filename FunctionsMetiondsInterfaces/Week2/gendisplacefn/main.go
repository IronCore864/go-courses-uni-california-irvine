package main

import (
	"fmt"
)

func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}

func main() {
	var a float64
	var v0 float64
	var s0 float64
	var t float64
	fmt.Println("Input ccceleration:")
	fmt.Scan(&a)
	fmt.Println("Input initial velocity:")
	fmt.Scan(&v0)
	fmt.Println("Input initial displacement:")
	fmt.Scan(&s0)
	fmt.Println("Input time:")
	fmt.Scan(&t)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println(fn(t))
}
