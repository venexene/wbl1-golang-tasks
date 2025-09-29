package main

import "fmt"



type Transport interface {
	Drive()
}

type Car struct {}

func (c *Car) Drive() {
	fmt.Println("Машина едет")
}



type Driver struct {
}

func (d *Driver) Travel(t Transport) {
	t.Drive()
}



type Animal interface {
	Move()
}

type Horse struct {}

func (h *Horse) Move() {
	fmt.Println("Лошадь скачет")
}



// Адаптер, который адаптирует структуру Horse к интерфейсу Transport
type AdapterHorseToCar struct {
	h *Horse
}

func (a *AdapterHorseToCar) Drive() {
	a.h.Move()
}



func main() {
	driver := &Driver{}
	car := &Car{}
	driver.Travel(car)

	horse := &Horse{}
	horseTransport := &AdapterHorseToCar{h: horse}
	driver.Travel(horseTransport)
}


/*
Применимость паттерна Адаптер:
- Если нужно использовать созданный ранее класс, а его интерфейс не совместим с новыми частями кода
- Если нужно реализовать совместную работу разных сторонних библиотек
- Если производится рефакторинг проекта и нужно сохранять работоспособность страхы компонентов

Плюсы:
- Скрывает детали преобразования интерфейса от клиентского кода
- Реализует повтороное использование кода без его изменения
- Может применяться для проведение мок-тестироваиня

Недостатки:
- Значительно загромождает и усложняет код
- Может снижать производительность, если используется в ресурсоемких задачах 
*/