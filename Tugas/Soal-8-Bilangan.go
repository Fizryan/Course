package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64

	fmt.Print("Masukkan Nilai X: ")
	fmt.Scan(&x)

	// Hitung nilai fungsi
	fungsiX := math.Sqrt((x*x*x + 3*x) * (x*x*x*x - 3*x*x + 4))

	fmt.Printf("Nilai f(x) adalah: %.2f\n", fungsiX)
}
