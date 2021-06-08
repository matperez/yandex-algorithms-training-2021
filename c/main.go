package main

import (
	"fmt"
	"regexp"
)

func main() {
	var a, b string
	fmt.Scan(&a)
	a = stripCountryCode(appendCode(stripNonDigits(a)))
	for i := 0; i < 3; i++ {
		fmt.Scan(&b)
		if stripCountryCode(appendCode(stripNonDigits(b))) == a {
			fmt.Print("YES\n")
		} else {
			fmt.Print("NO\n")
		}
	}
}

func stripCountryCode(num string) string {
	if len(num) == 11 {
		re := regexp.MustCompile("^7|8(\\d{10})$")
		return re.ReplaceAllString(num, "$1")
	}
	return num
}

func appendCode(num string) string {
	if len(num) == 7 {
		return "495" + num
	}
	return num
}

func stripNonDigits(num string) string {
	re := regexp.MustCompile("([^0-9])")
	return re.ReplaceAllString(num, "")
}