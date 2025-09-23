package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup


// Выход из горутины по условию
func stopOnCondition() {
	done := false

	wg.Add(1)
	go func() {
		count := 0
		defer wg.Done()

		for !done {
			count++
			fmt.Println("OnCondition: работа...")
			time.Sleep(1 * time.Second)

			if (count >= 5) {
				done = true
			}
		}

		fmt.Println("OnCondition: горутина завершена!")
	}()
}


// Выход из горутины по сигналу из канала
func stopWithCloseСhannel() {
	done := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <- done:
				fmt.Println("WithCloseChannel: горутина завершена!")
				return
			default:
				fmt.Println("WithCloseChannel: работа...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("WithCloseChannel: закрытие канала")
	close(done)
}


// Выход из горутины по закрытию канала
func stopWithSignalСhannel() {
	done := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <- done:
				fmt.Println("WithSignalChannel: горутина завершена!")
				return
			default:
				fmt.Println("WithSignalChannel: работа...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("WithSignalChannel: послан сигнал о завершении")
	done <- true
}


// Выход из горутины по завершению контекста
func stopWithContext() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <- ctx.Done():
				fmt.Println("WithContext: горутина завершена!")
				return
			default:
				fmt.Println("WithContext: работа...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()
}


// Выход из горутины по runtime.Goexit()
func stopOnGoExit() {
	wg.Add(1)
	go func() {
		count := 0
		defer wg.Done()
		defer fmt.Println("OnGoExit: горутина завершена!")

		for {
			count++
			fmt.Println("OnGoExit: работа...")
			time.Sleep(1 * time.Second)

			if (count >= 5) {
				runtime.Goexit()
			}
		}
	}()
}


func main() {
	stopOnGoExit()

	wg.Wait()
}