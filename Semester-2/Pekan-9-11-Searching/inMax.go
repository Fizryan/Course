package main

import "fmt"

const N int = 7

func main() {
	T := [N]int{11, -13, 1, 10, -11, 7, -9}
	awal, akhir := 2, 5
	fmt.Println(inMax(T, awal, akhir))
}

func inMax(T [N]int, awal int, akhir int) int {
	max := T[awal] // Inisialisasi nilai maksimum dengan nilai pada indeks awal

	for i := awal + 1; i <= akhir; i++ { // Mulai iterasi dari indeks awal+1 hingga akhir
		if T[i] > max { // Jika nilai pada indeks ke-i lebih besar dari nilai maksimum sebelumnya
			max = T[i] // Perbarui nilai maksimum
		}
	}

	return max
}
