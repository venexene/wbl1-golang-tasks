package main

import "fmt"


type Human struct {
	name string
	age int
}

func (h *Human) SayHello() {
	fmt.Printf("Hello, my name is %s. I am %d years old\n", h.name, h.age)
}


type Action struct {
	Human
	speed int
}

func (a *Action) Run () {
	fmt.Printf("I can run with speed of %d\n km/h", a.speed)
}


func main() {
	a := &Action{
		Human: Human{name: "Bob", age: 20},
		speed:12,
	}
	a.SayHello()
	a.Run()
}