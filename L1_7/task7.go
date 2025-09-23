package main

import (
	"fmt"
	"sync"
)

// Структура для map с мьютексом для безопсаной записи/чтения данных
type safeMap struct {
	elems map[int]string
	sync.RWMutex
}

// Запись в map с блокировкой мьютекса
func (sm *safeMap) set(key int, val string) {
	sm.Lock()
	defer sm.Unlock()

	sm.elems[key] = val
}

// Получение из map с блокировкой мьютекса на чтение
func (sm *safeMap) get(key int) (string, bool) {
	sm.RLock()
	defer sm.RUnlock()

	value, ok := sm.elems[key]
	return value, ok
}


func main() {
	sm := &safeMap{elems: make(map[int]string)}
	
	var wg sync.WaitGroup

	// Запись в map
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			value := fmt.Sprintf("value_%d", key)
			sm.set(key, value)
		}(i)
	}

	wg.Wait()

	// Одновременна запись и чтение 

	for i := 5; i < 10; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			value := fmt.Sprintf("value%d", key)
			sm.set(key, value)
		}(i)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			value, ok := sm.get(key)
			if !ok {
				fmt.Println("Нет такого ключа!")
				return
			}
			fmt.Printf("Key: %d; Value: %s\n", key, value)
		}(i)
	}

	wg.Wait()
}