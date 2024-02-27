package main

import (
	"fmt"
)

func main() {
	var lumin int
	fmt.Print("Input Koin Lumin = ")
	fmt.Scan(&lumin)

	quantum := 10
	galactum := 3
	nebula := 2
	stellaris := 20

	quantumcoins := lumin / (stellaris * nebula * galactum * quantum)
	lumin %= stellaris * nebula * galactum * quantum

	galactumcoins := lumin / (stellaris * nebula * galactum)
	lumin %= stellaris * nebula * galactum

	nebulacoins := lumin / (stellaris * nebula)
	lumin %= stellaris * nebula

	stellariscoins := lumin / stellaris
	lumin %= stellaris

	fmt.Printf("%d %d %d %d %d\n", quantumcoins, galactumcoins, nebulacoins, stellariscoins, lumin)
}
