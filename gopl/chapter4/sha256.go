package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"merorin.golang/gopl/chapter2/popcount"
)

const (
	SIZE = 32
)

func main() {
	// exercise 4.2
	mode := flag.String("mode", "SHA256", "Output mode")
	flag.Parse()
	println(*mode)
	switch *mode {
	case "SHA512":
		sum512()
	case "SHA384":
		sum384()
	default:
		sum256()
	}
}

// exercise 4.1
func countDifferentBits(c1, c2 *[SIZE]byte) int {
	count := 0
	for i := 0; i < SIZE; i++ {
		count += popcount.CountByNaiveLess(uint64(c1[i] ^ c2[i]))
	}
	return count
}

func sum256() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("result is %v\n", countDifferentBits(&c1, &c2))
}

func sum384() {
	c1 := sha512.Sum384([]byte("x"))
	c2 := sha512.Sum384([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func sum512() {
	c1 := sha512.Sum512([]byte("x"))
	c2 := sha512.Sum512([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
