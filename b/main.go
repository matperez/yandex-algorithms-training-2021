package main

import (
	"fmt"
)

func main() {
	var a, b, c int64
	fmt.Scan(&a, &b, &c)
	if check(a, b, c) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

func check(a, b, c int64) bool {
	return (a < b + c) && (b < a + c) && (c < a + b)
}