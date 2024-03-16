package main

import (
	"fmt"
)

func main() {
	var awalc, akhirc, step float64
	fmt.Scanf("%f %f %f", &awalc, &akhirc, &step)

	fmt.Printf("%10s %10s %10s %10s\n", "Celcius", "Reamur", "Fahrenheit", "Kelvin")
	for i := awalc; i <= akhirc; i += step {
		C := i
		R := reamur(C)
		F := fahrenheit(C)
		K := kelvin(C)
		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", C, R, F, K)
	}
}

func reamur(C float64) float64 {
	var r float64
	r = (4.0 / 5.0) * C
	return r
}

func fahrenheit(C float64) float64 {
	var F float64
	F = (9.0/5.0)*C + 32
	return F
}

func kelvin(C float64) float64 {
	var K float64
	K = C + 273
	return K
}
