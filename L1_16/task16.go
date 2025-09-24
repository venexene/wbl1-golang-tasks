package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr) / 2]
	
	left := []int{}
	middle := []int{}
	right := []int{}
	
	for _, x := range arr {
		if x < pivot {
			left = append(left, x)
		} else if x == pivot {
			middle = append(middle, x)
		} else if x > pivot {
			right = append(right, x)
		}
	}

	sortedLeft := quickSort(left)
	sortedRight := quickSort(right)

	result := append(sortedLeft, middle...)
	result = append(result, sortedRight...)
	return result
}

func main() {
	arr := []int{6, 7, 2, 5, 1}
	fmt.Println(quickSort(arr))
}