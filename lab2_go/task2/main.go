package main

import "fmt"

// CountCharacters подсчитывает количество вхождений каждого символа в строке.
func CountCharacters(s string) map[rune]int {
	counts := make(map[rune]int)

	for _, char := range s {
		counts[char]++
	}

	return counts
}

func main() {
	// Пример использования функции
	input := "hello world"
	result := CountCharacters(input)

	// Выводим результат
	for char, count := range result {
		fmt.Printf("%c: %d\n", char, count)
	}
}
