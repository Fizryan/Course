package main

import "fmt"

func main() {

	var jari float64 = 0
	fmt.Scan(&jari)
	fmt.Println("Masukkan Jari-Jari: ", jari)
	fmt.Print("Volume bola ", (4/3)*3.14*jari*jari*jari)
}