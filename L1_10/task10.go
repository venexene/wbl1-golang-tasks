package main

import "fmt"

func main() {
	temps := [8]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	slicedTemps := make(map[int][]float32) // Создание map срезов для результатов

	for _, temp := range temps {
		group := int(temp/10) * 10 // Вычисление группы 
		slicedTemps[group] = append(slicedTemps[group], temp) // Добавление в срез по групппе
	}

	for key, value := range slicedTemps {
		fmt.Printf("%d: %v\n", key, value)
	}
}