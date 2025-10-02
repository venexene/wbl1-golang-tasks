package main

import "fmt"

func deleteFromSlice(slice []int, ind int) []int {
	copy(slice[ind:], slice[ind+1:]) // Сдвиг элементов влево
	slice = slice[:len(slice)-1] // Уменьшение размера слайса
	return slice
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5}
	slice = deleteFromSlice(slice, 2)
	fmt.Println(slice)
}