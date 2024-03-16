package main

import "fmt"

func main() {

	var panjang, lebar, tinggi, berat float64
	var statusPaket bool

	fmt.Print("Masukkan Panjang: ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan Lebar: ")
	fmt.Scan(&lebar)
	fmt.Print("Masukkan Tinggi: ")
	fmt.Scan(&tinggi)
	fmt.Print("Masukkan Berat: ")
	fmt.Scan(&berat)
	fmt.Println(" ")

	var volumePaket float64 = panjang * lebar * tinggi

	var volumeMKubik float64 = volumePaket / 100.0

	var beratKg float64 = berat / 100.0

	fmt.Println("Volume m3: ", volumeMKubik)
	fmt.Println("Berat Paket: ", beratKg)

	if volumeMKubik < 1 && beratKg < 30 {
		statusPaket = true
		fmt.Print("Paket boleh dikirim. Status: ", statusPaket)
	} else {
		statusPaket = false
		fmt.Print("Paket tidak boleh dikirim. Status: ", statusPaket)
	}

}
