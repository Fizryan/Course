package main

import "fmt"

func main() {
	var berat1, berat2, berat3, resultGr, resultKg int
	fmt.Scan(&berat1)
	fmt.Println("Masukkan nilai berat pertama = ", berat1)
	fmt.Scan(&berat2)
	fmt.Println("Masukkan nilai berat kedua = ", berat2)
	fmt.Scan(&berat3)
	fmt.Println("Masukkan nilai berat ketiga = ", berat3)

	resultGr = berat1 + berat2 + berat3
	resultKg = resultGr / 1000

	if resultGr%1000 == 0 {
		fmt.Printf("Total berat belerang yaitu %d kg\n", resultKg)
	} else {
		fmt.Printf("Total berat belerang yaitu %d kg %d gram\n", resultKg, resultGr%1000)
	}
}