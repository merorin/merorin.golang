package main

import "fmt"

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // overflow

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // overflow

	var x uint8 = 1<<1 | 1<<5
	fmt.Printf("%08b\n", x)

	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)
}
