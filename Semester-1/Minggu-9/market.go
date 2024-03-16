package main

import "fmt"

func main() {
	// Input
	var totalBelanja int
	var member bool

	fmt.Print("Total belanja: ")
	fmt.Scan(&totalBelanja)

	fmt.Print("member? (true/false): ")
	fmt.Scan(&member)

	// Output
	fmt.Println("Membership?", member)
	fmt.Println("Diskon?", totalBelanja >= 100000)
	fmt.Println("Cashback?", totalBelanja >= 200000 && member == true)
}
