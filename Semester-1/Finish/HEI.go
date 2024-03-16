package main

import "fmt"

func main() {
	var a int
	var b string

	fmt.Print("masukkan bilang bulat: ")
	fmt.Scan(&a)

	fmt.Print("masukan kata: ")
	fmt.Scan(&b)

	for i := 0; i < a; i++ {
		fmt.Println(b)
	}
}
