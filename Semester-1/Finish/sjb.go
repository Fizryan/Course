package main

import (
	"fmt"
)

func sumOfDigits(x int) int {
	sum := 0

	for x > 0 {
		digit := x % 10
		sum += digit
		x /= 10
	}

	return sum
}

func main() {
	var input int
	fmt.Print("Masukkan bilangan bulat positif X: ")
	fmt.Scan(&input)

	result := sumOfDigits(input)

	fmt.Println("Hasil penjumlahan setiap digit pada", input, "adalah:", result)
}
