package main

import "fmt"

const N int = 1000

func main() {
	T1 := [N]int{2, 3, 5, 7, 11, 13, 17, 0, 0}  // Isi array sampai N, sisanya bisa diisi dengan 0
	T2 := [N]int{17, 13, 11, 7, 5, 3, 2, 0, 0}  // Isi array sampai N, sisanya bisa diisi dengan 0
	T3 := [N]int{2, 17, 3, 5, 7, 11, 17, 19, 0} // Isi array sampai N, sisanya bisa diisi dengan 0

	fmt.Println(yangTertukar(T1, 7))
	fmt.Println(yangTertukar(T2, 7))
	fmt.Println(yangTertukar(T3, 8))
}

func yangTertukar(T [N]int, total int) int {
	count := 0 // Inisialisasi jumlah nilai yang tidak terurut membesar

	for i := 1; i < total; i++ { // Mulai iterasi dari indeks kedua
		if T[i] < T[i-1] { // Jika nilai pada indeks ke-i lebih kecil dari nilai pada indeks sebelumnya
			count++ // Tambahkan jumlah nilai yang tidak terurut membesar
		}
	}

	return count
}
