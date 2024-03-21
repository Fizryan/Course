package main

import (
	"fmt"
)

func jumlah(a, b int, hasil *int) {
	/*
		{I.S terdefinisi bilangan bulat x dan y
		F.S hasil berisi jumlah x + y jika x ganjil dan y genap}
	*/
	if a%2 != 0 && b%2 != 1 {
		*hasil += a + b
	}
}

func main() {
	var a, b, hasil int
	fmt.Scan(&a, &b)
	for a != 0 || b != 0 {
		jumlah(a, b, &hasil)
		fmt.Scan(&a, &b)
	}
	fmt.Println(hasil)
}
