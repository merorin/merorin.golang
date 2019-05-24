package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if text == "end" {
			break
		}
		fmt.Println(tryParseFloatAndFormatWithComma(text))
	}
}

// exercise 3.11
func tryParseFloatAndFormatWithComma(s string) string {
	var signedIndex int
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		signedIndex = 1
	} else {
		signedIndex = 0
	}
	var decimalIndex = strings.Index(s, ".")

	if decimalIndex < 0 {
		return s[:signedIndex] + commaWithoutRecursive(s[signedIndex:])
	} else {
		return s[:signedIndex] + commaWithoutRecursive(s[signedIndex:decimalIndex]) + "." + commaForDecimal(s[decimalIndex+1:])
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

func commaForDecimal(s string) string {
	b := []byte(s)
	var buf bytes.Buffer

	for i := 0; i < len(b); i++ {
		if i > 0 && i%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteString(string(b[i]))
	}

	return buf.String()
}
