package main

import "fmt"

func main() {
	var jari float64
	fmt.Print("Masukkan Jari-Jari: ")
	fmt.Scan(&jari)

	// Menggunakan nilai pi sebagai 22/7
	pi := 22.0 / 7.0

	// Menghitung volume bola
	volume := (4.0 / 3.0) * pi * jari * jari * jari

	fmt.Printf("Volume bola: %.2f\n", volume)
}
