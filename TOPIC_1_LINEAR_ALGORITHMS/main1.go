package TOPIC_1_LINEAR_ALGORITHMS

import (
	"fmt"
	"math"
)

// 1. Угол задан в градусах, минутах и секундах.
// Найти его величину в радианах (с максимально возможной точностью).

func Main1() {
	fmt.Println("TOPIC_1_LINEAR_ALGORITHMS")
	fmt.Println("Task #1")

	//В одном радиане примерно 57 градусов, 17 минут и 45 секунд.

	//Объявление переменных: градусы, минуты, секунды.
	var degrees, minutes, seconds int

	fmt.Print("Введите градусы: ")
	fmt.Scan(&degrees)

	fmt.Print("Введите минуты: ")
	fmt.Scan(&minutes)

	fmt.Print("Введите секунды: ")
	fmt.Scan(&seconds)

	// Переводим всё в градусы и в тип float64, чтобы не терять точность:
	allInDegrees := float64(degrees) + float64(minutes)/60 + float64(seconds)/3600

	// Перевод в радианы:
	resultRadians := allInDegrees * math.Pi / 180

	// Так как результат должен быть с максимальной точностью то используем
	// константу из библиотеки math. Вот чему там равно это число:
	// Pi = 3.14159265358979323846264338327950288419716939937510582097494459

	//Выводим результат:
	fmt.Printf("\nУгол в радианах: %.15f.", resultRadians)

	fmt.Print("\n\n")
}
