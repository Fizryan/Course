package main

import "fmt"

func main() {
	var x, y float32
	var result float32

	fmt.Scan(&x, &y)
	result = 1/(3*(x*x)) + 10*y + 7
	fmt.Println(result)
}
