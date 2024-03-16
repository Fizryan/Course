package main

import "fmt"

func main() {
	var lumin int

	quantum := 0
	galactum := 0
	nebula := 0
	stellaris := 0

	fmt.Print("Input Koin Lumin : ")
	fmt.Scan(&lumin)

	for lumin >= 20 {
		stellaris++
		lumin -= 20
	}
	for stellaris >= 2 {
		nebula++
		stellaris -= 2
	}
	for nebula >= 3 {
		galactum++
		nebula -= 3
	}
	for galactum >= 10 {
		quantum++
		galactum -= 10
	}

	fmt.Println("Konversi Koin : ")
	fmt.Printf("%d %d %d %d %d\n", quantum, galactum, nebula, stellaris, lumin)
}
