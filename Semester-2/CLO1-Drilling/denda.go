package main

import (
	"fmt"
)

func denda(hariTerlambat int) {
	var denda int
	if hariTerlambat >= 0 && hariTerlambat <= 5 {
		denda = hariTerlambat * 1000
		fmt.Println(denda)
	} else if hariTerlambat >= 6 && hariTerlambat <= 10 {
		denda = hariTerlambat * 2000
		fmt.Println(denda)
	} else if hariTerlambat > 10 {
		fmt.Println("cabut keanggotaan")
	} else {
		fmt.Println("input bilangan bulat!! >= 0")
	}
}

func main() {
	var hariTerlambat int
	fmt.Scan(&hariTerlambat)
	denda(hariTerlambat)
}
