package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	set := make(map[int]bool) // Сет для элементов первого множества
	res := make(map[int]bool) // Сет для пересечений

	// Заполнение сета для дальнейшнего эффективного поиска пересечений
	for _, elem := range a {
		set[elem] = true
	}

	// Заполнение итогового сета
	for _, elem := range b {
		if set[elem] {
			res[elem] = true
		}
	}
	
	fmt.Print("{")
	for elem := range res {
		fmt.Printf(" %d ", elem)
	}
	fmt.Print("}\n")
}