package main

import (
	"fmt"
)

func jumlahDigit(N int) int {
	/*{ I.S. Terdefinisi sebuah bilangan bulat positif N.
	F.S. Fungsi mengembalikan jumlah digit dari bilangan N. }*/

	if N < 10 { // Ketika N kurang dari 10, itu hanya memiliki satu digit
		return 1
	} else {
		// Jumlah digit dari N adalah jumlah digit dari N/10 ditambah 1
		return 1 + jumlahDigit(N/10)
	}
}

func main() {
	var N int

	// Meminta input dari user
	fmt.Print("Masukan Nilai N: ")
	fmt.Scan(&N)

	// Memanggil fungsi rekursif jumlahDigit dan mencetak hasilnya
	fmt.Println("Jumlah digit dari bilangan", N, "adalah:", jumlahDigit(N))
}
