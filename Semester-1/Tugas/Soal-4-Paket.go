package main

import "fmt"

func main() {
	var panjang, lebar, tinggi, berat int
	var statusPaket bool

	fmt.Print("Masukkan Panjang (cm): ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan Lebar (cm): ")
	fmt.Scan(&lebar)
	fmt.Print("Masukkan Tinggi (cm): ")
	fmt.Scan(&tinggi)
	fmt.Print("Masukkan Berat (g): ")
	fmt.Scan(&berat)
	fmt.Println(" ")

	// Hitung volume dan berat dalam unit yang diinginkan
	volumeMKubik := float64(panjang*lebar*tinggi) / 1000000.0 // Convert cm ke m^3
	beratKg := float64(berat) / 1000.0                        // Convert gram ke kg

	fmt.Println("Volume m3: ", volumeMKubik)
	fmt.Println("Berat Paket (Kg): ", beratKg, "Kg /", berat, "g")

	// Periksa apakah paket dapat dikirim
	if volumeMKubik <= 2.0 && beratKg <= 2.0 {
		statusPaket = true
		fmt.Printf("Paket boleh dikirim. Status: %t\n", statusPaket)
	} else {
		statusPaket = false
		fmt.Printf("Paket tidak boleh dikirim. Status: %t\n", statusPaket)
	}
}
