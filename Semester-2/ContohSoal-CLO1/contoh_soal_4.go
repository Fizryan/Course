package main

import "fmt"

func Fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibo(n-2) + Fibo(n-1)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(Fibo(n))
}
