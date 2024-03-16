package main

import (
	"fmt"
	"math"
)

func isConsecutive(n int) bool {
	if n < 10 {
		return true
	}
	last := n % 10

	for n > 0 {
		prev := (n / 10) % 10

		if math.Abs(float64(last-prev)) != 1 {
			return false
		}

		last = prev
		n = n / 10
	}

	return true
}

func main() {
	fmt.Print("Masukkan sebuah bilangan bulat positif: ")
	var x int
	fmt.Scan(&x)

	fmt.Println("Bilangan", x, "adalah konsekutif:", isConsecutive(x))
}
