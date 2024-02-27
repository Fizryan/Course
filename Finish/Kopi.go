package main

import (
	"fmt"
)

func main() {
	var sugar, coffee, sugarreq, coffeereq uint
	fmt.Printf("Enter available sugar and coffee, then how much a cup needs\n")
	fmt.Print("Sugar: ")
	fmt.Scan(&sugar)
	fmt.Print("Coffee: ")
	fmt.Scan(&coffee)
	fmt.Print("SugarReq: ")
	fmt.Scan(&sugarreq)
	fmt.Print("CoffeeReq: ")
	fmt.Scan(&coffeereq)
	fmt.Printf("%d cups can be made", min(sugar/sugarreq, coffee/coffeereq))
}
