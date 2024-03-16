package main

import "fmt"

func main() {
	// Membaca jumlah anggota tim
	var anggotaTim int
	fmt.Print("Masukkan jumlah anggota tim: ")
	fmt.Scan(&anggotaTim)

	// Membaca status kelengkapan setiap anggota
	var result bool = true
	for i := 0; i < anggotaTim; i++ {
		var item1, item2, item3, item4, item5 bool

		// Membaca status kelengkapan 5 item dari setiap anggota
		fmt.Printf("Masukkan status kelengkapan item (true/false) untuk anggota %d: ", i+1)
		fmt.Scan(&item1, &item2, &item3, &item4, &item5)

		// Mengecek kelengkapan setiap anggota
		if !(item1 && item2 && item3 && item4 && item5) {
			result = false
			break
		}
	}

	// Menampilkan hasil kesiapan tim
	fmt.Println("Result:", result)
}
