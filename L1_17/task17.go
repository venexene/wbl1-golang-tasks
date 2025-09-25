package main

import "fmt"

func binarySearch(arr []int, target int, first int, last int) int {
	if last < first {
		return -1
	}

	middle := first + (last - first) / 2

	if arr[middle] == target {
		return middle
	} else if arr[middle] < target {
		return binarySearch(arr, target, middle + 1, last)
	} else {
		return binarySearch(arr, target, first, middle - 1)
	}
}

func main() {
	arr := []int{1, 3, 4, 5, 6, 7, 10, 14}
	target := 10

	pos := binarySearch(arr, target, 0, len(arr) - 1)
	fmt.Printf("Позиция искомого: %d\n", pos)
}