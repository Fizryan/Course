package main

import (
	"fmt"
)

type Rectangle struct {
	panjang, lebar int
	warna          string
}

type Geometry struct {
	Area, Perimeter int
}

func main() {
	var rect Rectangle
	isiData(&rect)
	hitung(&rect)
	fmt.Printf("Luas: %d\n", rect.panjang)
	fmt.Printf("Keliling: %d\n", rect.lebar)
}

func isiData(persegi *Rectangle) {
	fmt.Print("Masukkan panjang dan lebar persegi panjang: ")
	fmt.Scan(&persegi.panjang, &persegi.lebar)
	fmt.Print("Masukkan warna persegi panjang: ")
	fmt.Scan(&persegi.warna)
}

func hitung(persegi *Rectangle) {
	persegi.panjang = persegi.panjang * persegi.lebar
	persegi.lebar = 2 * (persegi.panjang + persegi.lebar)
}
