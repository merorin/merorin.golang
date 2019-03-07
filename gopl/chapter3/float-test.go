package main

import (
	"fmt"
	"math"
)

func main() {
	// 因为float32的有效bit位只有23个，其它的bit位用于指数和符号；当整数大于23bit能表达的范围时，float32的表示将出现误差
	var f float32 = 1 << 24
	fmt.Println(f == (f + 1)) // it will be true

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.4f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
}
