package main

import "fmt"

func jumlahA(kata string, index int) int {
	count := 0 // Variabel untuk menyimpan jumlah huruf 'a'

	// Loop untuk memeriksa setiap karakter dalam kata
	for i := 0; i < len(kata); i++ {
		if kata[i] == 'a' {
			count++ // Menambahkan jumlah huruf 'a' jika ditemukan
		}
	}

	return count // Mengembalikan jumlah huruf 'a'
}

func main() {
	var x string
	var result, index int
	index = 0
	fmt.Scanln(&x)
	result = jumlahA(x, index)
	fmt.Println(result)
}
