package main

import "fmt"

// Структура Human
type Human struct {
	name string
	age int
}

func (h *Human) SayHello() {
	fmt.Printf("Hello, my name is %s. I am %d years old\n", h.name, h.age)
}


// Структура Action, которая применяет встраивание структуры Human
type Action struct {
	Human
	speed int
}

func (a *Action) Run () {
	fmt.Printf("I can run with speed of %d\n km/h", a.speed)
}


func main() {
	// Создание экземпляра структуры Action
	a := &Action{
		// Инициалиазция встраиваемой структуры Human
		Human: Human{name: "Bob", age: 20},
		speed: 12,
	}
	a.SayHello() // Вызов метода из встроенной структуры Human
	a.Run() // Вызов собственного метода структуры Action
}