package main

import (
	"fmt"
	"merorin.golang/gopl/chapter2/popcount"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	value, _ := strconv.ParseInt(arg, 0, 64)
	fmt.Printf("result for %d is %d\n", value, popcount.CountByTable(uint64(value)))
	fmt.Printf("result for %d is %d\n", value, popcount.CountByLoopTable(uint64(value)))
	fmt.Printf("result for %d is %d\n", value, popcount.CountByShift(uint64(value)))
	fmt.Printf("result for %d is %d\n", value, popcount.CountByNaiveLess(uint64(value)))
}
