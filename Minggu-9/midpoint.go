package main

import "fmt"

func numtengah(a, b, c int) bool {
	// Urutkan tiga bilangan
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}

	// Cek apakah b adalah nilai tengah
	return a < b && b < c
}

func main() {
	// Input
	var num1, num2, num3 int
	fmt.Println("Masukkan tiga bilangan bulat:")
	fmt.Scan(&num1, &num2, &num3)

	// Output
	result := numtengah(num1, num2, num3)
	fmt.Println(result)
}
