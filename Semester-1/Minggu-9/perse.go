package main

import "fmt"

func main() {
	var panjang, lebar int

	// Minta pengguna memasukkan panjang dan lebar persegi panjang
	fmt.Print("Masukkan panjang persegi panjang: ")
	fmt.Scan(&panjang)

	fmt.Print("Masukkan lebar persegi panjang: ")
	fmt.Scan(&lebar)

	// Hitung keliling persegi panjang
	keliling := 2 * (panjang + lebar)

	// Tampilkan hasil
	fmt.Printf("Keliling persegi panjang: %d\n", keliling)
}
