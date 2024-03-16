package main

import "fmt"

func main() {
	var bilangan int

	// Input bilangan dari pengguna
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&bilangan)

	// Cek apakah bilangan merupakan kelipatan 3 dan/atau kelipatan 5
	if bilangan%3 == 0 && bilangan%5 == 0 {
		fmt.Println("Kelipatan 3 dan Kelipatan 5")
	} else if bilangan%3 == 0 {
		fmt.Println("Kelipatan 3")
	} else if bilangan%5 == 0 {
		fmt.Println("Kelipatan 5")
	} else {
		fmt.Println("")
	}
}
