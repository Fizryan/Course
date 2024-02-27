package main

import "fmt"

func Fizryan(tahun int) bool {
	return (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0)
}

func main() {
	// Input
	var inputTahun int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&inputTahun)

	// Output
	result := Fizryan(inputTahun)
	fmt.Println(result)
}
