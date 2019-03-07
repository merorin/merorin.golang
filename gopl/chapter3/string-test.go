package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, world"

	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	//c := s[len(s)] // panic: runtime error: index out of range
	//fmt.Println(c)

	fmt.Println(s[0:5])

	fmt.Println(s[:5])

	fmt.Println("goodbye" + s[5:])

	s = "left foot"
	t := s
	s += ", right foot"

	fmt.Println(s)
	fmt.Println(t)

	s = "Hello, 世界"
	fmt.Println(s)
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeLastRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	n := 0
	for range s {
		n++
	}
	fmt.Println(n)

	s = "プログラム"
	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("% x\n", r)
}
