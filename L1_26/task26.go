package main

import (
	"fmt"
	"strings"
)

func checkUniqueness(str string) bool {
	str = strings.ToLower(str)
	set := make(map[rune]bool)

	for _, char := range str {
		if set[char] {
			return false
		}
		set[char] = true
	}

	return true
}

func main() {
	str := "abcd"
	fmt.Printf("\"%s\" -> %t\n", str,  checkUniqueness(str))

	str = "abCdefAaf"
	fmt.Printf("\"%s\" -> %t\n", str, checkUniqueness(str))

	str = "aabcd"
	fmt.Printf("\"%s\" -> %t\n", str, checkUniqueness(str))
}