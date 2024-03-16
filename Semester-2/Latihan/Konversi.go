package main

import (
	"fmt"
)

func main() {
	var bilangan int
	var hasil string
	fmt.Scan(&bilangan)
	hasil = numToTeks(bilangan)
	fmt.Println(hasil)
}

func numToTeks(bilangan int) string {
	var konversi string
	if bilangan < 0 {
		konversi = "negatif"
	} else if bilangan > 0 {
		konversi = "positif"
	} else {
		konversi = "netral"
	}
	return konversi

}
