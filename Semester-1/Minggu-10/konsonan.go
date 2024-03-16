package main

import (
	"fmt"
	"strings"
)

func main() {
	var huruf string
	fmt.Print("Masukkan satu huruf: ")
	fmt.Scan(&huruf)

	huruf = strings.ToLower(huruf)

	// Cek apakah huruf konsonan atau bukan
	if isKonsonan(huruf) {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan konsonan")
	}
}

// Fungsi untuk mengecek apakah huruf konsonan atau bukan
func isKonsonan(huruf string) bool {

	daftarKonsonan := "bcdfghjklmnpqrstvwxyz"

	return strings.Contains(daftarKonsonan, huruf)
}
