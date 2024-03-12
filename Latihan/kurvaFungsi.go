package main

import "fmt"

// Prosedur KurvaFungsiKuadrat untuk menentukan arah kurva fungsi kuadrat
func KurvaFungsiKuadrat(a, b, c float64) {
	if a > 0 {
		fmt.Println("terbuka ke atas")
	} else {
		fmt.Println("terbuka ke bawah")
	}
}

func main() {
	var a, b, c float64

	// Meminta pengguna memasukkan koefisien a, b, dan c
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	// Memanggil prosedur KurvaFungsiKuadrat
	KurvaFungsiKuadrat(a, b, c)
}
