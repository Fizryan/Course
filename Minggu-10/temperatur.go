package main

import (
	"fmt"
)

func main() {
	var temperatures [5]float64

	fmt.Println("Masukkan lima catatan temperatur:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Pencatatan %d: ", i+1)
		fmt.Scan(&temperatures[i])
	}

	status := getStatus(temperatures)

	fmt.Println("Result:", status)
}

func getStatus(temperatures [5]float64) string {
	naik := false
	turun := false

	for i := 1; i < 5; i++ {
		if temperatures[i] > temperatures[i-1] {
			naik = true
		} else if temperatures[i] < temperatures[i-1] {
			turun = true
		}
	}

	if naik && turun {
		return "Temperatur Tidak stabil"
	} else if naik {
		return "Temperatur Stabil naik"
	} else if turun {
		return "Temperatur Stabil turun"
	}

	return "Tidak stabil"
}
