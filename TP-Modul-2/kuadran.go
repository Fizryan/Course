package main

import (
	"fmt"
)

func main() {
	var x, y float64

	fmt.Println("Masukkan Titik Koordinat x dan y:")
	fmt.Print("x: ")
	fmt.Scan(&x, &y)
	fmt.Print("y: ")
	fmt.Scan(&y)

	var kuadran string
	if x > 0 && y > 0 {
		kuadran = "Kuadran I"
	} else if x < 0 && y > 0 {
		kuadran = "Kuadran II"
	} else if x < 0 && y < 0 {
		kuadran = "Kuadran III"
	} else if x > 0 && y < 0 {
		kuadran = "Kuadran IV"
	} else if x == 0 && y == 0 {
		kuadran = "Center"
	} else if x == 0 && y != 0 {
		kuadran = "Titik Y"
	} else if x != 0 && y == 0 {
		kuadran = "Titik X"
	} else {
		kuadran = "Titik"
	}

	fmt.Println("Pesawat Musuh Berada Di Posisi ", kuadran)
}
