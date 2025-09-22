package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
	"os/signal"
	"syscall"
)


// Функция для отправки в канал
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

// Функция для чтения из канала
func getFromChannel(ctx context.Context, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case n, ok := <- c:
			if !ok { // Завершение при закрытии канала
				return
			}
			fmt.Printf("Получено: %s\n", n)
		case <-ctx.Done(): // Завршение при отмене контекста
			return
		}
	}
}

func main() {
	// Проверка на то, что при вызове был указан аргумент
	if len(os.Args) < 2 {
		fmt.Println("Укажите время работы!")
		return
	}

	// Преобразование аргумента
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Некорректное время работы!")
		return
	}

	// Создание контекста с таймаутом
	timeCtx, timeCancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer timeCancel() // Определение освобождения ресурсов по таймауту

	// Создание контекста, который отменяется при Ctrl+C
	ctx, cancel := signal.NotifyContext(timeCtx, syscall.SIGINT)
	defer cancel() // Определение остановки отслеживания сигнала

	// Создание канала
	ch := make(chan string)

	// Создание WaitGroup для ожидания завершения гоуртин
	var wg sync.WaitGroup

	// Запуск горутины для чтения из канала
	wg.Add(1)
	go getFromChannel(ctx, ch, &wg)	

	// Запуск горутины для чтения ввода и записи в канал
	go sendToChannel(ctx, ch)

	// Ожидание отмены контекста 
	<-ctx.Done()

	// Определение причины завершения
	select {
	case <-timeCtx.Done():
		fmt.Println("Время работы истекло!")
	case <-ctx.Done():
		fmt.Println("Работа завершена вручную досрочно!")
	}

	close(ch) // Закрытие канала 
	wg.Wait() // Ожидание заврешения горутин
}