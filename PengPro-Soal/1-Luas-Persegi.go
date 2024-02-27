package main

import "fmt"

func main() { 
    var panjang, lebar, result int

    // Jangan Lupa Masukan Nilai
	fmt.Scan(&panjang)
	fmt.Println("Masukan Nilai Panjang: ", panjang)
	fmt.Scan(&lebar)
	fmt.Println("Masukan Nilai Lebar: ", lebar)

	result = panjang*lebar
	fmt.Println("Luas Persegi Panjang Adalah = ", result)
}