package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func reverseString(str string) string {
	runes := []rune(str)

	for i := 0; i < len(runes) / 2; i++ {
		runes[i], runes[(len(runes) - 1) - i] = runes[(len(runes) - 1) - i], runes[i]
	}

	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	str, _ := reader.ReadString('\n')
	str = strings.TrimRight(str, "\r\n") // Удаление лишнего переноса строки

	reversedStr := reverseString(str)
	fmt.Printf("Перевернутая строка: %s\n", reversedStr)
}