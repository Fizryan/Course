package main

import "fmt"

func main() {
	var input, total int

	total = 0
	fmt.Println("Nilai Input: ")

	for {
		fmt.Scan(&input)

		if input%2 != 0 {
			break
		}

		total = total + input
	}

	fmt.Println("Nilai Output:", total)
}
