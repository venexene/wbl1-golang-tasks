package main

import "fmt"

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool) // Множество уникальных элементов

	// Заполнение множества
	for _, elem := range a {
		set[elem] = true
	}

	// Вывод содержимого множества
	fmt.Print("{")
	for elem := range set {
		fmt.Printf(" \"%s\" ", elem)
	}
	fmt.Print("}\n")
}