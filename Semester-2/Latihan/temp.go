package main

import (
	"fmt"
)

// KonversiSuhu merupakan prosedur untuk mengkonversi suhu dari Celsius ke Reaumur, Fahrenheit, dan Kelvin
func KonversiSuhu(celsius float64) (reaumur, fahrenheit, kelvin float64) {
	reaumur = celsius * 0.8
	fahrenheit = (celsius * 9 / 5) + 32
	kelvin = celsius + 273.15
	return reaumur, fahrenheit, kelvin
}

func main() {
	var celsius float64
	fmt.Scanln(&celsius)

	reaumur, fahrenheit, kelvin := KonversiSuhu(celsius)

	fmt.Printf("%.2f %.2f %.2f\n", reaumur, fahrenheit, kelvin)
}
