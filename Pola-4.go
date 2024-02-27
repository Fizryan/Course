package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukkan bilangan bulat positif x: ")
	fmt.Scan(&x)

	if x <= 0 {
		fmt.Println("Masukan harus merupakan bilangan bulat positif.")
		return
	}

	fmt.Println("Keluaran:")
	for i := 1; i <= x; i++ {
		if i == 1 || i == x {
			for j := 1; j <= x; j++ {
				fmt.Print(i, " ")
			}
		} else {
			fmt.Print(i)
			fmt.Print("           ")
			fmt.Print(i)
		}
		fmt.Println()
	}
}
