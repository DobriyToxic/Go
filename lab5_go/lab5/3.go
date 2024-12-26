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

// filterParallel обрабатывает одну строку изображения, преобразуя её в оттенки серого.
// Параметры:
// - img: изображение для обработки (интерфейс draw.RGBA64Image позволяет изменять пиксели).
// - wg: указатель на sync.WaitGroup для синхронизации завершения горутин.
// - rowIndex: индекс строки, которую нужно обработать.
func filterParallel(img draw.RGBA64Image, wg *sync.WaitGroup, rowIndex int) {
	defer wg.Done() // Уменьшаем счётчик WaitGroup после завершения горутины.

	// Получаем ширину изображения.
	bounds := img.Bounds()
	width := bounds.Max.X

	// Обрабатываем каждый пиксель в строке.
	for x := 0; x < width; x++ {
		// Получаем цвет текущего пикселя.
		rgba := img.RGBA64At(x, rowIndex)

		// Вычисляем среднее значение (оттенок серого) на основе компонентов R, G, B.
		gray := uint16((uint32(rgba.R) + uint32(rgba.G) + uint32(rgba.B)) / 3)

		// Устанавливаем новый цвет пикселя (оттенок серого).
		img.SetRGBA64(x, rowIndex, color.RGBA64{R: gray, G: gray, B: gray, A: rgba.A})
	}
}

func main() {
	// Получаем текущую рабочую директорию для информации.
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка при получении текущей директории:", err)
		return
	}
	fmt.Println("Текущая рабочая директория:", dir)

	// Открываем файл с исходным изображением.
	file, err := os.Open("input.png")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close() // Закрываем файл после использования.

	// Декодируем изображение из файла.
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Ошибка при декодировании изображения:", err)
		return
	}

	// Приводим изображение к интерфейсу draw.RGBA64Image, чтобы можно было изменять пиксели.
	drawImg, ok := img.(draw.RGBA64Image)
	if !ok {
		fmt.Println("Ошибка: изображение не может быть преобразовано в draw.RGBA64Image")
		return
	}

	// Начинаем замер времени обработки.
	startTime := time.Now()

	// Получаем размеры изображения (высота нужна для обработки строк).
	bounds := drawImg.Bounds()
	height := bounds.Max.Y

	// Создаем объект sync.WaitGroup для синхронизации всех горутин.
	var wg sync.WaitGroup

	// Запускаем горутины для обработки каждой строки изображения.
	for y := 0; y < height; y++ {
		wg.Add(1)                          // Увеличиваем счётчик WaitGroup перед запуском горутины.
		go filterParallel(drawImg, &wg, y) // Обрабатываем строку y в отдельной горутине.
	}

	// Ждём завершения всех горутин.
	wg.Wait()

	// Выводим время, затраченное на обработку изображения.
	fmt.Println("Время обработки с параллельными горутинами:", time.Since(startTime))

	// Создаём файл для сохранения обработанного изображения.
	outputFile, err := os.Create("output_parallel.png")
	if err != nil {
		fmt.Println("Ошибка при создании файла для сохранения:", err)
		return
	}
	defer outputFile.Close() // Закрываем файл после использования.

	// Сохраняем обработанное изображение в формате PNG.
	err = png.Encode(outputFile, drawImg)
	if err != nil {
		fmt.Println("Ошибка при сохранении изображения:", err)
		return
	}

	// Уведомляем пользователя об успешном завершении.
	fmt.Println("Изображение успешно обработано и сохранено в output_parallel.png")
}
