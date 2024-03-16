package main

import (
	"fmt"
)

func main() {
	// Langkah 1: Masukkan nomor voucher belanja
	var x int
	fmt.Print("Masukkan nomor voucher belanja (4 digit): ")
	fmt.Scan(&x)

	// Langkah 2: Cek apakah dua digit di tengah adalah genap
	digitTengah := (x / 10) % 100
	genap := digitTengah%2 == 0

	// Langkah 3: Cek apakah penjumlahan digit pertama dan ketiga adalah kelipatan digit keempat
	digitPertama := x / 1000
	digitKetiga := (x / 10) % 10
	digitKeempat := x % 10

	// Periksa apakah digitKeempat adalah 0 untuk menghindari pembagian oleh nol
	kelipatan := digitKeempat != 0 && (digitPertama+digitKetiga)%digitKeempat == 0

	// Langkah 4: Tampilkan hasil diskon dan cashback
	fmt.Println("Diskon?", genap)
	fmt.Println("Cashback?", kelipatan)
}
