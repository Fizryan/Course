package main

import "fmt"

func main() {
	var N, i int
	var berhenti bool
	fmt.Scan(&N)
	i = 0
	berhenti = false
	for !berhenti {
		i = i + 1
		fmt.Print(i, ",")
		berhenti = i == N
	}
}
