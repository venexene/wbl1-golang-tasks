package main

import (
	"fmt"
	"sync"
)

type counter struct {
	val int
	sync.RWMutex
}

func (c *counter) increment() {
	c.Lock()
	defer c.Unlock()
	c.val++
}

func (c *counter) get() int {
	c.RLock()
	defer c.RUnlock()
	return c.val
}

func main() {
	c := &counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Итоговое значение счетчика: %d\n", c.get())
}