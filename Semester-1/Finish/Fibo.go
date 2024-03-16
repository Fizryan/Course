package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("N harus lebih besar dari 1")
		return
	}

	fibonacci := make([]int, n+1)
	fibonacci[0] = 0
	fibonacci[1] = 1

	for i := 2; i <= n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}

	fmt.Println("Deret Fibonacci hingga suku ke-", n, ":")
	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci[i], " ")
	}
	fmt.Println()
}
