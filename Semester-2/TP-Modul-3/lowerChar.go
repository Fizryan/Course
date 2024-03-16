package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	hasil := lowToUpper(input)
	fmt.Println(hasil)
}

func lowToUpper(kar string) string {
	result := ""
	for _, char := range kar {
		if char >= 'a' && char <= 'z' {
			result += string(char - 32)
		} else {
			result += string(char)
		}
	}
	return result
}
