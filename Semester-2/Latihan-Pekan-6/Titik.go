package main

import (
	"fmt"
	"math"
)

type titik struct {
	x float64
	y float64
}

func jarak(p1, p2 titik) float64 {
	a := p1.x - p2.x
	b := p1.y - p2.y
	return akar(a*a + b*b)
}

func akar(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	var p1, p2 titik

	fmt.Print("Masukkan koordinat titik pertama (x1 y1): ")
	fmt.Scan(&p1.x, &p1.y)

	fmt.Print("Masukkan koordinat titik kedua (x2 y2): ")
	fmt.Scan(&p2.x, &p2.y)
	titik1 := titik{p1.x, p1.y}
	titik2 := titik{p2.x, p2.y}

	fmt.Printf("Jarak antara dua titik adalah: %.2f\n", jarak(titik1, titik2))
}
