package main

import (
	"fmt"
)

func main() {
	var b1, b2, hasilJumlah int
	fmt.Scanln(&b1, &b2)
	hasilJumlah = penjumlahan(b1, b2)
	fmt.Println(hasilJumlah)
}

func penjumlahan(b1, b2 int) int {
	perhitungan = (b1 + b2)
	return perhitungan
}
