package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	initCh := make(chan int)
	doubleCh := make(chan int)

	// Запись чисел из массива в первый канал
	wg.Add(1)
	go func(ch chan<- int) {
		defer wg.Done()
		for _, x := range arr{
			ch <- x
		}
		close(ch)
	}(initCh)

	// Получение из первого канала и запись удвоенного числа во второй канал
	wg.Add(1)
	go func(readCh <-chan int, writeCh chan<- int) {
		defer wg.Done()
		for x := range readCh {
			writeCh <- x * 2
		}
		close(writeCh)
	}(initCh, doubleCh)

	// Вывод чисел из второго канала
	wg.Add(1)
	go func (ch <-chan int) {
		defer wg.Done()
		for x := range ch {
			fmt.Println(x)
		}
	}(doubleCh)

	wg.Wait()
}