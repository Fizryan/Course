package main

import "fmt"

func main() {
	var n, luas, sisi, keliling int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&sisi)

		luas = sisi * sisi
		keliling = 4 * sisi

		fmt.Printf("Luas: %d\nKeliling %d")
	}
}
