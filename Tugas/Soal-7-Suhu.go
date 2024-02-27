package main

import "fmt"

func main() {

	var celcius int
	fmt.Print("Masukan nilai temperature dalam Celcius: ")
	fmt.Scan(&celcius)
	fmt.Println("Nilai temperatur dalam Kelvin adalah: ", celcius+273)
}
