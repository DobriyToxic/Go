package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// Установка режима release для продакшн-среды
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Настройка доверенных прокси (замените IP на ваши настройки)
	router.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	// Маршрут для сложения
	router.GET("/add", calculateHandler(func(a, b int) int {
		return a + b
	}))

	// Маршрут для вычитания
	router.GET("/sub", calculateHandler(func(a, b int) int {
		return a - b
	}))

	// Маршрут для умножения
	router.GET("/mul", calculateHandler(func(a, b int) int {
		return a * b
	}))

	// Маршрут для деления
	router.GET("/div", func(c *gin.Context) {
		a, b, err := getQueryParams(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if b == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Деление на ноль невозможно"})
			return
		}
		result := a / b
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	// Запуск сервера на другом порту
	if err := router.Run(":8081"); err != nil {
		panic(err)
	}
}

// calculateHandler создаёт обработчик для арифметических операций
func calculateHandler(operation func(a, b int) int) gin.HandlerFunc {
	return func(c *gin.Context) {
		a, b, err := getQueryParams(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := operation(a, b)
		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}

// Функция для извлечения и преобразования query-параметров a и b
func getQueryParams(c *gin.Context) (int, int, error) {
	aStr := c.Query("a")
	bStr := c.Query("b")

	if aStr == "" || bStr == "" {
		return 0, 0, http.ErrMissingFile
	}

	a, err := strconv.Atoi(aStr)
	if err != nil {
		return 0, 0, err
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		return 0, 0, err
	}

	return a, b, nil
}
