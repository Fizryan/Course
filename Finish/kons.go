package main

import (
	"fmt"
)

func isConsecutive(x int) bool {
	prevDigit := -1

	for x > 0 {
		digit := x % 10

		if prevDigit != -1 && abs(prevDigit-digit) != 1 {
			return false
		}

		prevDigit = digit
		x /= 10
	}

	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	var input int
	fmt.Print("Masukkan bilangan bulat positif X: ")
	fmt.Scan(&input)

	result := isConsecutive(input)

	if result {
		fmt.Println(input, "adalah bilangan konsekutif.")
	} else {
		fmt.Println(input, "bukan bilangan konsekutif.")
	}
}
