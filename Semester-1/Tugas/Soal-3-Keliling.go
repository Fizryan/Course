package main

import "fmt"

func main() {

	var panjang, lebar, result int
	fmt.Print("Masukan Panjang: ")
	fmt.Scan(&panjang)
	fmt.Print("Masukan Lebar: ")
	fmt.Scan(&lebar)
	result = panjang + lebar*2
	fmt.Println("Keliling persegi panjang adalah= ", result)
}
