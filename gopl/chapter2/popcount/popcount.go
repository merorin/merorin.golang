package popcount

var pc [256]byte

//var pc = func() (pc [256]byte) {
//	for i := range pc {
//		pc[i] = pc[i / 2] + byte(i&1)
//	}
//	return
//}()

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func CountByTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// exercise 2.3
func CountByLoopTable(x uint64) int {
	var result byte = 0
	for i := 0; i < 8; i++ {
		result += pc[byte(x>>uint(i*8))]
	}
	return int(result)
}

// exercise 2.4
func CountByShift(x uint64) int {
	count := 0
	for ; x > 0; x >>= 1 {
		if x&1 == 1 {
			count++
		}
	}
	return count
}

// exercise 2.5
func CountByNaiveLess(x uint64) int {
	count := 0
	for ; x > 0; count++ {
		x &= x - 1
	}
	return count
}
