package main

import "fmt"

func main() {
    
    var panjang, lebar, result int
	fmt.Scan(&panjang)
	fmt.Println("Masukan Nilai Panjang: ", panjang)
    fmt.Scan(&lebar)
    fmt.Println("Masukan Nilai Lebar: ", lebar)
    result = panjang+lebar*2
	fmt.Println("Keliling persegi panjang adalah= ", result)
}