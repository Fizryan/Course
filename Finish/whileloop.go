package main

// while-loop program
import "fmt"

func main() {
	var input, total int

	// jika ganjil maka hasilnya 0
	total = 0

	fmt.Println("Input:")
	fmt.Scan(&input)

	// cek genap
	for input%2 == 0 {
		total = total + input

		// minta input hingga ganjil
		fmt.Scan(&input)
	}
	fmt.Println("Output:", total)
}
