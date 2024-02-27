package main

import (
	"fmt"
	"strconv"
	"strings"
)

func analisisGol(gol []int) (int, int) {
	// Mencari jumlah gol terbanyak dan terendah
	golTerbanyak := gol[0]
	golTerendah := gol[0]

	for _, nilai := range gol {
		if nilai > golTerbanyak {
			golTerbanyak = nilai
		}
		if nilai < golTerendah {
			golTerendah = nilai
		}
	}

	return golTerbanyak, golTerendah
}

func main() {
	// Input dari pengguna
	fmt.Print("Masukkan jumlah gol dalam format 4 angka terpisah: ")
	var inputGol string
	fmt.Scanln(&inputGol)

	// Memisahkan string input menjadi array integer
	inputGolArray := make([]int, 4)
	inputGolSlice := strings.Fields(inputGol)
	for i, nilai := range inputGolSlice {
		gol, err := strconv.Atoi(nilai)
		if err != nil {
			fmt.Println("Error dalam mengonversi input.")
			return
		}
		inputGolArray[i] = gol
	}

	// Analisis jumlah gol
	golTerbanyak, golTerendah := analisisGol(inputGolArray)

	// Output
	fmt.Printf("Jumlah gol terbanyak dan terendah: %d %d\n", golTerbanyak, golTerendah)
}
