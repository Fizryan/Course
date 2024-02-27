package main

import "fmt"

func main() {
	var x, result float64
    fmt.Scan(&x)
	fmt.Println("Masukkan nilai x = ", x)
	result = ((x * x * x) + (3 * x)) / ((x * x * x * x) - (3 * x * x) + 4)
	fmt.Printf("Hasil nilai y adalah = %v", result)
}