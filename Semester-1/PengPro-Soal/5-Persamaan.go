package main

import "fmt"

func main() {
	var x, result float64
	fmt.Scan(&x)
	fmt.Println("Masukkan Nilai X = ", x)
	result = (3*x - 5) * (2*x + 1)
	fmt.Println("Nilai Y adalah= ", result)
}