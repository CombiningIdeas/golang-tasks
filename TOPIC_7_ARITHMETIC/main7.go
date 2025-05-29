package TOPIC_7_ARITHMETIC

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1. Натуральное число в p-ичной системе счисления задано своими цифрами,
// хранящимися в массиве K(n). Проверить корректность такого представления и
// перевести число в q-ичную систему (возможно, число слишком велико, чтобы
// получить его внутреннее представление; кроме того, p<=10, q<=10).

func Main7() {
	fmt.Println("TOPIC_7_ARITHMETIC")
	fmt.Println("Task #1")

	// считываем исходные данные систем счисления:
	var p, q, lenDigit int64
	fmt.Print("Введите p (основание исходной системы счисления, <=10): ")
	fmt.Scan(&p)
	fmt.Print("Введите q (основание итоговой системы счисления, <=10): ")
	fmt.Scan(&q)
	fmt.Print("Введите количество цифр числа: ")
	fmt.Scan(&lenDigit)

	if p < 2 || p > 10 || q < 2 || q > 10 || lenDigit <= 0 {
		fmt.Println("Некорректные значения p, q или кол-ва цифр числа")
		return
	}

	// Проверка на то, не больше ли чем нужно пользователь ввел цифр с клавиатуры
	// Это нужно, чтобы избежать визуальной ошибки с консоли и при проверке данного алгоритма
	fmt.Printf("Введите %d цифр числа в системе счисления %d "+
		"(в одну строку, через пробел):\n", lenDigit, p)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	// Проверяем ввод пользователя
	if int64(len(parts)) != lenDigit {
		fmt.Println("Ошибка: количество введённых цифр не соответствует ожидаемому.")
		return
	}

	// Создаем нужный массив K и заполняем его
	K := make([]int64, lenDigit)
	for iter, symbol := range parts {
		num, err := strconv.ParseInt(symbol, 10, 64)
		if err != nil || num < 0 || num >= p {
			fmt.Printf("Ошибка: цифра \"%s\" некорректна для основания %d\n",
				symbol, p)
			return
		}
		K[iter] = num // Записываем в массив текущую цифру исходного числа
	}

	// Перевод в десятичную систему счисления:
	var decimalNumberSystem int64 = 0
	for _, iter := range K {
		decimalNumberSystem = decimalNumberSystem*p + iter
	}

	// Проверяем не равно ли число 0, иначе выводи результат
	if decimalNumberSystem == 0 {
		fmt.Println("Результат: 0")
		return
	}

	// Перевод в q-ичную систему счисления:
	// Объявляем срез для хранения результата — цифр числа в системе счисления q
	var resultArray []int64

	// Пока десятичное число больше нуля — продолжаем делить его на q
	for decimalNumberSystem > 0 {
		// Получаем остатот от деления, это младшая цифра в итоговой системе счисления:
		remainderOfDivision := decimalNumberSystem % q
		// Добавляем эту цифру в начало итогового массива, от старших к младшим цифрам
		resultArray = append([]int64{remainderOfDivision}, resultArray...)
		// Делим число на q, чтобы перейти к следующей, более старшей цифре
		decimalNumberSystem /= q
	}

	fmt.Print("Результат: ")
	for _, resultDigits := range resultArray {
		fmt.Print(resultDigits)
	}

	fmt.Printf("\n\n")
}
