package main

import "fmt"

func main() {
	var x, result float64

	fmt.Print("Masukkan Nilai X: ")
	fmt.Scan(&x)

	result = 6*x*x + 3*x - 10*x - 5

	fmt.Printf("Nilai Y adalah: %.2f\n", result)
}
