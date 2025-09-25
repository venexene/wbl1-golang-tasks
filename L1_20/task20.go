package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func reverseString(runes []rune, left int, right int) {
	for left < right {
        runes[left], runes[right] = runes[right], runes[left]
        left++
        right--
    }
}

func reverseWords(str string) string {
	runes := []rune(str)
	
	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			reverseString(runes, start, i-1)
			start = i + 1
		}
	}
	
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	str, _ := reader.ReadString('\n')
    str = strings.TrimSpace(str)

	reversedStr := reverseWords(str)
	fmt.Printf("Строка с перевернутыми словами: %s\n", reversedStr)
}