package main

import "fmt"

func denominisasi(amount int) (int, int, int) {
	amount10k := 0
	amount2k := 0
	amount1k := 0

	for amount >= 10000 {
		amount10k++
		amount -= 10000
	}

	for amount >= 2000 {
		amount2k++
		amount -= 2000
	}

	for amount >= 1000 {
		amount1k++
		amount -= 1000
	}
	return amount10k, amount2k, amount1k
}

func main() {
	var amount int
	amount10k := 0
	amount2k := 0
	amount1k := 0
	fmt.Scan(&amount)

	if amount >= 0 {
		amount10k, amount2k, amount1k = denominisasi(amount)
		fmt.Printf("%d lembar 10000\n", amount10k)
		fmt.Printf("%d lembar 2000\n", amount2k)
		fmt.Printf("%d lembar 1000\n", amount1k)
	} else {
		fmt.Println("Enter Valid Amount")
	}
}
