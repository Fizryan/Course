package main

import (
	"fmt"
)

func uangHotel(kota string, hari int) int {
	var tarif int

	//menentukan tarif tiap kota
	switch kota {
	case "A":
		tarif = 500000
	case "B":
		tarif = 400000
	case "C":
		tarif = 350000
	}
	//menghitung total
	total := (hari - 1) * tarif
	return total
}

func main() {
	var kota string
	var hari, total int
	fmt.Scan(&kota, &hari)
	total = uangHotel(kota, hari)
	fmt.Println(total)
}
