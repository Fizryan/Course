package main

import "fmt"

// Prosedur konsekutif untuk menentukan apakah suatu bilangan adalah konsekutif atau tidak
func konsekutif(bilangan int) {
	digitTerakhir := bilangan % 10
	bilangan /= 10

	for bilangan > 0 {
		digit := bilangan % 10
		if digit != digitTerakhir+1 && digit != digitTerakhir-1 {
			fmt.Println(false)
			return
		}
		digitTerakhir = digit
		bilangan /= 10
	}

	fmt.Println(true)
}

func main() {
	var bilangan int

	// Meminta pengguna memasukkan bilangan bulat positif
	fmt.Scan(&bilangan)

	// Memanggil prosedur konsekutif
	konsekutif(bilangan)
}
