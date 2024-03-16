package main

import "fmt"

// Fungsi untuk menukar nilai dua variabel
func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

// Fungsi untuk mengurutkan dua bilangan secara membesar
func sort(x *int, y *int) {
	if *x > *y {
		swap(x, y)
	}
}

func main() {
	var x, y int

	// Membaca masukan sampai x dan y bernilai sama
	for {
		fmt.Scan(&x, &y)
		if x == y {
			break
		}

		// Panggil fungsi untuk mengurutkan x dan y
		sort(&x, &y)

		// Menampilkan nilai x dan y setelah diurutkan secara membesar
		fmt.Println(x, y)
	}
}
