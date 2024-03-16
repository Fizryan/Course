package main

import (
	"fmt"
	"strconv"
)

func main() {

	var userInput, digit1, digit2 int
	var d1String, d2String string // d1 itu maksudnya digit 1, dan d2 itu digit 2

	fmt.Print("Masukan Bilangan 2 digit: ")
	_, err := fmt.Scan(&userInput)

	if err != nil {
		fmt.Println("Input yang anda masukkan bukan bilangan bulat!!")
	} else if userInput < 100 && userInput > 9 {
		digit1 = (userInput / 10)
		d1String = strconv.Itoa(digit1)
		digit2 = (userInput % 10)
		d2String = strconv.Itoa(digit2)
		fmt.Println("> ", d1String, d1String, d2String, d2String)
	} else {
		fmt.Println("Nilai yang anda masukkan tidak memenuhi syarat!!")
		fmt.Println("Masukan Nilai angka > 9 dan < 100")
	}
}
