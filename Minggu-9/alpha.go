package main

import (
	"fmt"
	"unicode"
)

func Alpha(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsNumber(char)
}

func main() {
	// Input
	var input rune
	fmt.Printf("Masukkan karakter: ")
	fmt.Scanf("%c", &input)

	// Output
	result := Alpha(input)
	fmt.Println(result)
}
