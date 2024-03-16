package main

import (
	"fmt"
)

func main() {
	var rA, rB, jarak float64 // r itu jari-jari
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&rA, &rB, &jarak)

	// Menentukan apakah lingkaran A dan B beririsan atau tidak
	result := irisan(rA, rB, jarak)

	// Menampilkan hasil
	fmt.Println(result)
}

// Menentukan apakah dua lingkaran beririsan atau tidak
func irisan(rA, rB, jarak float64) bool {
	// Menghitung jumlah jari-jari
	sum := rA + rB

	// Memeriksa apakah jarak antara titik pusat kurang dari atau sama dengan jumlah jari-jari
	if jarak <= sum {
		return false
	} else {
		return true
	}
}
