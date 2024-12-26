package main

import (
	"errors"
	"fmt"
)

// FormatIP принимает массив из 4 байтов и возвращает строку в формате IP-адреса.
func FormatIP(ip [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// ListEven принимает диапазон (два целых числа) и возвращает срез четных чисел и ошибку.
func ListEven(start, end int) ([]int, error) {
	if start > end {
		return nil, errors.New("левая граница больше правой")
	}

	var evens []int
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}

	return evens, nil
}

func main() {
	// Пример работы FormatIP
	ip := [4]byte{192, 168, 0, 1}
	fmt.Println("IP-адрес:", FormatIP(ip))

	// Пример работы ListEven
	start, end := 1, 10
	evens, err := ListEven(start, end)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Чётные числа:", evens)
	}

	// Пример с ошибкой
	evens, err = ListEven(10, 1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
