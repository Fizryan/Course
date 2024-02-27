package main

import "fmt"

func segitiga(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

func main() {
	// Input
	var a, b, c int
	fmt.Println("Masukkan tiga bilangan bulat:")
	fmt.Scan(&a, &b, &c)

	// Output
	result := segitiga(a, b, c)
	fmt.Println(result)
}
