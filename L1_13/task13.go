package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Printf("a = %d, b = %d\n", a, b)

	a = a ^ b // Применение XOR здесь дает сумму a и b
	b = a ^ b // Применение XOR здесь дает занчение а
	a = a ^ b // Применение XOR здесь дает занчение b
	fmt.Printf("a = %d, b = %d\n", a, b)
}