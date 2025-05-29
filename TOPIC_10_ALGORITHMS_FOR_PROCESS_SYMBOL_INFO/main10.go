package TOPIC_10_ALGORITHMS_FOR_PROCESS_SYMBOL_INFO

import (
	"bufio"
	"fmt"
	"os"
)

// 1. Текст записан одной длинной строкой. Признаком красной строки служит символ $.
// Переформатировать текст в 60-символьные строки, формируя абзацы.

// Функция форматирования текста
func formatTextCharByChar(input string) {
	const lineWidth int = 60  // Предел длины строки после которого будет красный абзац
	indentation := "    "     // Красная строка, это 4 пробела
	var symbolCounter int = 0 // Счётчик символов в текущей строке
	newParagraph := true      // Флаг отвечающий за то, нужна ли красная строка

	for _, ch := range input {
		if ch == '$' {
			// Абзац окончен — перенос строки
			fmt.Println()
			newParagraph = true
			symbolCounter = 0
			continue
		}

		// Вставляем отступ, если начинается новый абзац
		if newParagraph {
			fmt.Print(indentation)
			symbolCounter += len(indentation)
			newParagraph = false
		}

		fmt.Printf("%c", ch)
		symbolCounter++

		// Если достигнута длина строки — перенос строки
		if symbolCounter == lineWidth {
			fmt.Println()
			symbolCounter = 0
			newParagraph = true
		}
	}
	// Завершающий перенос строки
	if symbolCounter > 0 {
		fmt.Println()
	}
}

func Main10() {
	fmt.Println("TOPIC_10_ALGORITHMS_FOR_PROCESS_SYMBOL_INFO")
	fmt.Println("Task #1")

	// Считываем строку введённую с консоли:
	fmt.Println("Введите длинную строку(можете ввести разделить $):")

	// Тут наша строка, хранящаяся в сканере буфера:
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

	// Проверка на то, что пользователь что-то вел и строка не пустая:
	if scanner.Scan() {
		var input string = scanner.Text() // Получили считанную строку в строковом виде
		fmt.Printf("\nОтформатированный текст: \n\n")
		formatTextCharByChar(input) //Выводим результат
	}
}
