package main

import (
	"fmt"
)

func kuadran(x, y int) {
	/*
		I.S. terdefinisi dua buah bilangan bulat positif atau negatif yang menyatakan nilai x dan y dari sebuah titik.
		F.S. keluarkan bilangan bulat 1, 2, 3, atau 4 jika masing-masing berada pada kuadran 1, 2, 3, atau 4.
	*/
	if x > 0 && y > 0 {
		fmt.Println("1")
	} else if x < 0 && y > 0 {
		fmt.Println("2")
	} else if x < 0 && y < 0 {
		fmt.Println("3")
	} else if x > 0 && y < 0 {
		fmt.Println("4")
	} else {
		fmt.Println("0")
	}
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	kuadran(x, y)
}
