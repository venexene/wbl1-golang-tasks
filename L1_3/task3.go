package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"bufio"
)


// Воркер для получения данных из канала
func worker(id int, c chan string, wg *sync.WaitGroup) {
	// Объявление уменьшения счетчика по завершению фунцкции
	defer wg.Done()

	// Чтение данных из канала до его закрытия
	for n := range c {
		fmt.Printf("Воркер %d получил: %s\n", id, n)
	}
}


func main() {
	// Проверка на то, что при вызове был указан аргумент
	if len(os.Args) < 2 {
		fmt.Println("Укажите количество воркеров!")
		return
	}

	// Преобразование аргумента
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Некорректное количество воркеров!")
		return
	}

	// Создание канала
	ch := make(chan string)

	// Создание WaitGroup для ожидания завершения гоуртин
	var wg sync.WaitGroup

	// Запуск полученного числа воркеров
	for i := 0; i < n; i++ {
		wg.Add(1) // Увеличение счетчика ожидания
		go worker(i, ch, &wg) // Запуск воркера в горутине
	}

	// Создание сканера для чтения ввода
	scanner := bufio.NewScanner(os.Stdin)

	// Чтение ввода
	for scanner.Scan() {
		text := scanner.Text()
		ch <- text
	}

	close(ch) // Закрытие канала 
	wg.Wait() // Ожидание заврешения горутин
}