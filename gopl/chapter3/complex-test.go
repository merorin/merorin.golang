package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var x = complex(1, 2) // equals to "x := 1 + 2i"
	var y = complex(3, 4) // equals to "y := 3 + 4i"

	fmt.Println(x * y) // (a + bi) * (c + di) = (ac - bd) + (bc + ad)i = -5 + 10i
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))

	fmt.Println(cmplx.Sqrt(-1))
}
