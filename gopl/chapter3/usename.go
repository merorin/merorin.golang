package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expertBasename("a/b/c.go")) // "c"
	fmt.Println(expertBasename("c.d.go"))   // "c.d"
	fmt.Println(expertBasename("abc"))      // "abc"

	fmt.Println(advancedBasename("a/b/c.go")) // "c"
	fmt.Println(advancedBasename("c.d.go"))   // "c.d"
	fmt.Println(advancedBasename("abc"))      // "abc"
}

func expertBasename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func advancedBasename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
