package main

import (
	"fmt"
)

func main() {
	var Start, speedAli, speedBadu, jarakAli, jarakBadu int

	Start = (8 * 3600)
	speedAli = 5
	speedBadu = 8
	jarakAli = 0
	jarakBadu = 0
	for jarakBadu <= jarakAli {
		jarakAli += speedAli
		if Start >= 32400 {
			jarakBadu += speedBadu
		}
		Start++
		if jarakBadu >= jarakAli {
			break
		}
	}
	fmt.Printf("Titik Ali %02dm titik Badu %02dm\n", jarakAli, jarakBadu)
	fmt.Printf("Badu mendahului Ali pada pukul %02d:%02d:%02d\n", (Start / 3600), (Start%3600)/60, (Start % 60))
}
