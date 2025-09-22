package main

import (
	"fmt"
	"sync"
)


func main() {
	arr := []int{2, 4, 6, 8, 10}

	// Создание WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	for elem := range arr {
		// Увеличение счетчика на 1 перед запуском новой горутины
		wg.Add(1)

		// Запус горутины
		go func(n int) {
			// Объявление уменьшения счетчика по завершению фунцкции
			defer wg.Done()
			fmt.Println(n*n) // Вывод квадарата чисел 
		}(elem) // Передача текущего элемента среза в горутину
	}

	wg.Wait() // Ожидание завершения горутин
}