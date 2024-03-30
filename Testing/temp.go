package main

import (
	"fmt"
)

// KonversiSuhu merupakan prosedur untuk mengkonversi suhu dari Celsius ke Reaumur, Fahrenheit, dan Kelvin
func KonversiSuhu(celsius float64, step int) (reaumur, fahrenheit, kelvin float64) {
	if step == 0 {
		return celsius * 0.8, (celsius * 9 / 5) + 32, celsius + 273.15
	}

	reaumur, fahrenheit, kelvin = KonversiSuhu(celsius, step-1)

	switch step {
	case 3:
		kelvin = (fahrenheit-32)*5/9 + 273.15
	case 2:
		fahrenheit = (reaumur*1.25 + 32)
	case 1:
		reaumur = celsius * 0.8
	}

	return reaumur, fahrenheit, kelvin
}

func main() {
	var celsius float64
	fmt.Scanln(&celsius)

	reaumur, fahrenheit, kelvin := KonversiSuhu(celsius, 3)

	fmt.Printf("%.2f %.2f %.2f\n", reaumur, fahrenheit, kelvin)
}
