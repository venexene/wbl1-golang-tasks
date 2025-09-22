package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)


// Функция для отправки в канал с отменой через контекст
func sendToChannel(ctx context.Context, c chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		select {
		case c <- scanner.Text(): // Отправка данных в канал
		case <-ctx.Done(): // Выход из функции при отмене контекста
			return
		}
	}
}


// Воркер для получения данных из канала с отменой через контекст
func worker(ctx context.Context, id int, c chan string, wg *sync.WaitGroup) {
	// Объявление уменьшения счетчика по завершению фунцкции
	defer wg.Done() 

	for {
		select {
		case n, ok := <- c:
			if !ok { // Завершение при закрытии канала
				return
			}
			fmt.Printf("Воркер %d получил: %s\n", id, n)
		case <-ctx.Done(): // Завршение при отмене контекста
			return
		}
	}
}


func main() {
	// Создание контекста, который отменяется при Ctrl+C
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop() // Определение остановки отслеживания сигнала

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
		go worker(ctx, i, ch, &wg) // Запуск воркера в горутине
	}

	// Создание сканера для чтения ввода
	go sendToChannel(ctx, ch)

	// Ожидание отмены контекста
	<-ctx.Done()

	close(ch) // Закрытие канала 
	wg.Wait() // Ожидание заврешения горутин
}