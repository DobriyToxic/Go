package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
	"time"
)

// filter выполняет преобразование части изображения в оттенки серого.
// Она обрабатывает строки изображения от startY до endY.
// Параметры:
// - img: изображение для обработки.
// - wg: объект WaitGroup для синхронизации горутин.
// - startY, endY: диапазон строк, который обрабатывается этой функцией.
// - width: ширина изображения.
func filter(img draw.RGBA64Image, wg *sync.WaitGroup, startY, endY, width int) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения работы горутины.

	// Проходим по каждой строке и колонке в заданном диапазоне.
	for y := startY; y < endY; y++ {
		for x := 0; x < width; x++ {
			// Получаем цвет текущего пикселя.
			rgba := img.RGBA64At(x, y)

			// Вычисляем среднее значение (интенсивность) цветов RGB для преобразования в оттенки серого.
			gray := uint16((uint32(rgba.R) + uint32(rgba.G) + uint32(rgba.B)) / 3)

			// Устанавливаем новый цвет пикселя (оттенок серого).
			img.SetRGBA64(x, y, color.RGBA64{R: gray, G: gray, B: gray, A: rgba.A})
		}
	}
}

func main() {
	// Получаем текущую рабочую директорию.
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка при получении текущей директории:", err)
		return
	}
	fmt.Println("Текущая рабочая директория:", dir)

	// Открываем входной файл с изображением.
	file, err := os.Open("input.png")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close() // Закрываем файл после завершения работы.

	// Декодируем изображение из файла.
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Ошибка при декодировании изображения:", err)
		return
	}

	// Приводим изображение к типу draw.RGBA64Image для возможности изменения пикселей.
	drawImg, ok := img.(draw.RGBA64Image)
	if !ok {
		fmt.Println("Ошибка: изображение не может быть преобразовано в draw.RGBA64Image")
		return
	}

	// Начинаем отсчет времени обработки.
	startTime := time.Now()

	// Получаем размеры изображения.
	bounds := drawImg.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Создаем объект WaitGroup для синхронизации горутин.
	var wg sync.WaitGroup

	// Задаем количество горутин для параллельной обработки.
	numGoroutines := 16
	rowsPerGoroutine := height / numGoroutines // Число строк на одну горутину.

	// Запускаем горутины для обработки частей изображения.
	for i := 0; i < numGoroutines; i++ {
		startY := i * rowsPerGoroutine    // Начальная строка для этой горутины.
		endY := startY + rowsPerGoroutine // Конечная строка для этой горутины.
		if i == numGoroutines-1 {         // Последняя горутина обрабатывает оставшиеся строки.
			endY = height
		}

		wg.Add(1)                                    // Увеличиваем счетчик WaitGroup.
		go filter(drawImg, &wg, startY, endY, width) // Запускаем горутину для обработки.
	}

	// Ожидаем завершения всех горутин.
	wg.Wait()

	// Выводим время, затраченное на обработку.
	fmt.Println("Время обработки:", time.Since(startTime))

	// Создаем файл для сохранения обработанного изображения.
	outputFile, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Ошибка при создании файла для сохранения:", err)
		return
	}
	defer outputFile.Close() // Закрываем файл после завершения работы.

	// Кодируем обработанное изображение в файл PNG.
	err = png.Encode(outputFile, drawImg)
	if err != nil {
		fmt.Println("Ошибка при сохранении изображения:", err)
		return
	}

	// Уведомляем пользователя об успешном завершении.
	fmt.Println("Изображение успешно обработано и сохранено в output.png")
}
