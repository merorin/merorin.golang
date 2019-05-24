package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Input two strings split by space...")
	for input.Scan() {
		text := input.Text()
		if text == "end" {
			break
		}
		stringArray := strings.Fields(text)
		if len(stringArray) < 2 {
			fmt.Println("Insufficient strings")
		} else {
			fmt.Println(compare(stringArray[0], stringArray[1]))
		}
	}
}

func compare(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return false
	}

	arr1 := strings.Split(s1, "")
	arr2 := strings.Split(s2, "")

	sort.Strings(arr1)
	sort.Strings(arr2)

	// O(nlogn) + O(n)
	return strings.Join(arr1, "") == strings.Join(arr2, "")
}
