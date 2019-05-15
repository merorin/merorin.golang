package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if text == "end" {
			break
		}
		fmt.Println(commaWithoutRecursive(text))
	}
}

// exercise 3.10
func commaWithoutRecursive(s string) string {
	n := len(s)
	var buf bytes.Buffer
	firstCommaIndex := n % 3 // first comma index
	if firstCommaIndex <= 0 {
		firstCommaIndex += 3
	}
	buf.WriteString(s[:firstCommaIndex])

	for i := firstCommaIndex; i < n; {
		buf.WriteString(",")
		var nextIndex int
		if (i + 3) >= n {
			nextIndex = n
		} else {
			nextIndex = i + 3
		}
		buf.WriteString(s[i:nextIndex])
		i = nextIndex
	}

	return buf.String()
}
