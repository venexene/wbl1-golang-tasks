package main

import (
	"fmt"
	"reflect"
)

func printType(x interface{}) {
	tp := ""

	// switch по возможным типам x
	switch x.(type) {
	case int:
		tp = "int"
	case string:
		tp = "string"
	case bool:
		tp = "bool"
	default:
		// Определения типа chan с помощью рефлексии
		if reflect.TypeOf(x).Kind() == reflect.Chan {
			tp = "chan"
		} else {
			tp = "unknown type"
		}
	}
	fmt.Printf("Тип переменной: %s\n", tp)
}

func main() {
	i := 1
	printType(i)

	str := "hello"
	printType(str)

	b := true
	printType(b)

	ch := make(chan int)
	printType(ch)
}