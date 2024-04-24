package main

import "fmt"

const N int = 1000

func main() {
	var T [N]int
	var total, bilX int

	fmt.Scan(&total)
	for i := 0; i < total; i++ {
		fmt.Scan(&T[i])
	}
	fmt.Scan(&bilX)
	fmt.Print(indeks(T, total, bilX))
}

func indeks(T [N]int, total, bilX int) int {
	/*
		algoritma untuk mencari indeks dari bilX, jika tidak ditemukan akan mengembalikan -1
		mengembalikan indeks dari bilX
	*/
	for i := 0; i < total; i++ {
		if T[i] == bilX {
			return i
		}
	}
	return -1
}
