package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari, tinggi float64

	fmt.Print("Masukkan Jari-Jari Penampang Silinder: ")
	fmt.Scan(&jariJari)

	fmt.Print("Masukkan Tinggi Silinder: ")
	fmt.Scan(&tinggi)

	luas := 2 * math.Pi * jariJari * (jariJari + tinggi)

	fmt.Printf("Luas Silinder: %.2f\n", luas)
}
