package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Создаем новый роутер Gin
	r := gin.Default()

	// Обрабатываем GET-запрос с query-параметрами name и age
	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")

		// Формируем ответ
		response := "Меня зовут " + name + ", мне " + age + " лет"

		// Отправляем ответ в виде строки
		c.String(200, response)
	})

	// Запускаем сервер на порту 8080
	r.Run(":8080")
}
