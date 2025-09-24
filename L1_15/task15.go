package main

// Проблема была в неэффективном использовании памяти

func createHugeString(size int) string {
	bytes := make([]byte, size)
	for i := range bytes {
		bytes[i] = 'x'
	}
	return string(bytes)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString := make([]byte, 100)
	copy(justString, v[:100]) // Копирование данных, чтобы не ссылаться на часть большой строки
}

func main() {
	someFunc()
}