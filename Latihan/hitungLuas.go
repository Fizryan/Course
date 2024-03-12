package main

import "fmt"

// Prosedur hitungLuas untuk menghitung luas persegi panjang
func hitungLuas(p, l int, luas *int) {
	*luas = p * l
}

func main() {
	var panjang, lebar, luasPersegiPanjang int

	// Meminta pengguna memasukkan panjang dan lebar persegi panjang
	fmt.Scan(&panjang, &lebar)

	// Memanggil prosedur hitungLuas
	hitungLuas(panjang, lebar, &luasPersegiPanjang)

	// Menampilkan luas persegi panjang
	fmt.Println(luasPersegiPanjang)
}
