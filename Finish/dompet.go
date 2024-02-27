package main

import "fmt"

func main() {
	var uang int
	dompet := 0

	for {
		fmt.Print("uang: ")
		fmt.Scan(&uang)

		if uang != 0 {
			dompet += uang
		} else {
			break
		}
	}

	fmt.Println("Total uang di dompet:", dompet)
}
