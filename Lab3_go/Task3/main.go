package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Маршрут для задания 3: Подсчет символов
	router.POST("/count_chars", func(c *gin.Context) {
		// Структура для хранения входящих данных
		var jsonData struct {
			Text string `json:"text"`
		}

		// Прочитаем JSON-данные из запроса
		if err := c.BindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильный формат JSON"})
			return
		}

		// Подсчет символов
		charCount := make(map[rune]int)
		for _, char := range jsonData.Text {
			charCount[char]++
		}

		// Возвращаем результат в формате JSON
		c.JSON(http.StatusOK, charCount)
	})

	router.Run(":8082")
}
