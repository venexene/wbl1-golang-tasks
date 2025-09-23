package main

import "fmt"

func main() {
	var num int64
	var i int
	
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Printf("Ошибка ввода: %f", err)
		return
	}

	fmt.Print("Введите индекс бита: ")
	_, err = fmt.Scan(&i)
	if err != nil {
		fmt.Printf("Ошибка ввода: %f", err)
		return
	}

	// Сдвиг бита на i и применение исключающего ИЛИ для смены бита
	new_num := num ^ (1 << i)

	fmt.Printf("До: %d\n", num)
	fmt.Printf("После: %d\n", new_num)
}