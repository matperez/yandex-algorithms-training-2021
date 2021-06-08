package main

import (
	"fmt"
)

func main() {
	var num1, num2 int64
	var mode string
	fmt.Scan(&num1, &num2)
	fmt.Scan(&mode)
	switch mode {
	case "fan":
		fmt.Print(num1)
	case "heat":
		if num2 >= num1 {
			fmt.Print(num2)
		} else {
			fmt.Print(num1)
		}
	case "freeze":
		if num2 <= num1 {
			fmt.Print(num2)
		} else {
			fmt.Print(num1)
		}
	case "auto":
		fmt.Print(num2)
	}
}