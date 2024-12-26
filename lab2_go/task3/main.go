package main

import (
	"fmt"
	"math"
)

// Структура "Точка"
type Point struct {
	X, Y float64
}

// Структура "Отрезок"
type Segment struct {
	Start, End Point
}

// Структура "Треугольник"
type Triangle struct {
	A, B, C Point
}

// Структура "Круг"
type Circle struct {
	Center Point
	Radius float64
}

// Метод для вычисления длины отрезка
func (s Segment) Length() float64 {
	return math.Sqrt(math.Pow(s.End.X-s.Start.X, 2) + math.Pow(s.End.Y-s.Start.Y, 2))
}

// Метод для вычисления площади треугольника
func (t Triangle) Area() float64 {
	return math.Abs(0.5 * (t.A.X*(t.B.Y-t.C.Y) + t.B.X*(t.C.Y-t.A.Y) + t.C.X*(t.A.Y-t.B.Y)))
}

// Метод для вычисления площади круга
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Интерфейс Shape с методом Area
type Shape interface {
	Area() float64
}

// Функция, которая принимает тип Shape и выводит площадь фигуры
func printArea(s Shape) {
	result := s.Area()
	fmt.Printf("Площадь фигуры: %.2f\n", result)
}

func main() {
	// Пример для отрезка
	segment := Segment{Start: Point{0, 0}, End: Point{3, 4}}
	fmt.Printf("Длина отрезка: %.2f\n", segment.Length())

	// Пример для треугольника
	triangle := Triangle{
		A: Point{0, 0},
		B: Point{4, 0},
		C: Point{0, 3},
	}
	printArea(triangle)

	// Пример для круга
	circle := Circle{
		Center: Point{0, 0},
		Radius: 5,
	}
	printArea(circle)
}
