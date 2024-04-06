package main

import "fmt"

const NMax int = 1000

type tabT [NMax]int

func main() {
	var n, i int
	var DM, M, x, y tabT

	// Input n
	fmt.Scan(&n)

	// Input DM, M, x, y
	for i = 0; i < n; i++ {
		fmt.Scan(&DM[i], &M[i], &x[i], &y[i])
	}

	// Menampilkan Informasi banyaknya ramuan yang dapat dibuat
	for i = 0; i < n; i++ {
		// Hitung banyaknya ramuan yang dapat dibuat
		ramuan := min(DM[i]/x[i], M[i]/y[i])
		fmt.Printf("total ramuan yang berhasil dibuat dari lemari persediaan %d sebanyak: %d\n", i+1, ramuan)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
