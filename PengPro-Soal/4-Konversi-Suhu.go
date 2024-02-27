package main

import "fmt"

func main() {

	var celcius int

	fmt.Scan(&celcius)
	fmt.Println("Masukan nilai temperature dalam Celcius: ", celcius)
	fmt.Println("Nilai temperatur dalam Kelvin adalah: ", celcius+273)
}
