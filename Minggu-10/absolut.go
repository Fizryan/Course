package main

import (
	"fmt"
	"math"
)

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bilangan)

	// Hitung nilai mutlak menggunakan fungsi Abs dari package math
	Absolut := int(math.Abs(float64(bilangan)))

	// Tampilkan hasil
	fmt.Printf("Nilai mutlak dari %d adalah %d\n", bilangan, Absolut)
}
