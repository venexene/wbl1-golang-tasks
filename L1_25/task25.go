package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	done := make(chan struct{}) // Канал для опвещении о завершении ожидания
	go func() {
		defer close(done)
		start := time.Now() // Получение времени начала ожидания
		for time.Since(start) < duration {
			// Ожидание в течении duration 
		}
	}()
	<-done
}

func main() {
	fmt.Println("Сон на 5 секунд запущен!")
	sleep(5 * time.Second)
	fmt.Println("Сон завершен!")
}