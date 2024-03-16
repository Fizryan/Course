package main

import "fmt"

func main() {
	var T, N int
	fmt.Scan(&T) // Membaca kapasitas tanki
	fmt.Scan(&N) // Membaca jumlah ember yang akan diisi ke tanki

	var tanki, volume int
	results := make([]bool, 0)

	for i := 0; i < N; i++ {
		fmt.Scan(&volume) // Membaca volume ember
		tanki += volume   // Tambahkan volume ember ke tanki

		if tanki >= T { // Cek apakah tanki sudah penuh
			results = append(results, true)
			tanki = T // Jika penuh, isi tanki sampai penuh
		} else {
			results = append(results, false)
		}
	}

	// Menampilkan hasil
	for _, isFull := range results {
		fmt.Println(isFull)
	}
}
