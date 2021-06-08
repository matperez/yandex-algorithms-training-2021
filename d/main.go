package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	if a == 0 {
		if int64(math.Sqrt(float64(b))) == c {
			fmt.Print("MANY SOLUTIONS")
		} else {
			fmt.Print("NO SOLUTION")
		}
		return
	}
	x := (c * c - b)/a
	if check(x, a, b, c) {
		fmt.Print(x)
		return
	}
	fmt.Print("NO SOLUTION")
}

func check(x, a, b, c int64) bool {
	return int64(math.Sqrt(float64(a * x + b))) == c
}