package TOPIC_3_CYCLIC_AND_ITERATIVE_ALGORITHMS

import (
	"fmt"
	"math"
)

// 1. Численно убедиться, является ли заданная функция y=f(x) чётной или нёчетной
// на заданном отрезке -a <= x <= a. Учесть погрешность вычислений
// и возможные точки разрыва функции.

// Функция, которую проверяем, по дефолту это просто x, можно задать свою
func givenFunction(x float64) float64 {
	return x
}

// Функция для сравнение с учётом погрешности
func almostEqual(a, b, epsilon float64) bool {
	// Возвращает true, если числа почти равны, с точностью до epsilon
	// Если разность меньше epsilon, значит числа достаточно близки, и считаются равными.
	return math.Abs(a-b) < epsilon

}

func Main3() {
	fmt.Println("TOPIC_3_CYCLIC_AND_ITERATIVE_ALGORITHMS")
	fmt.Println("Task #1")

	var a float64
	var steps int
	const epsilon float64 = 1e-9 // Погрешность сравнения

	fmt.Print("Введите значение a (границы отрезка [-a; a]): ")
	fmt.Scan(&a)

	// Этот параметр будет управлять числом точек, в которых будет производиться
	// проверка функции на чётность/нечётность:
	fmt.Print("Введите количество шагов проверки: ")
	fmt.Scan(&steps)

	if a <= 0 || steps <= 0 {
		fmt.Println("Ошибка: 'a' и 'steps' должны быть положительными.")
		return
	}

	// Вычисление размера одного шага на отрезке от -a до a:
	stepSize := (2 * a) / float64(steps)

	// Это флаги, которые предполагают, что функция по умолчанию и чётная, и нечётная.
	// Они будут сбрасываться на false, если найдётся хотя бы одна точка,
	// где функция нарушает признак чётности или нечётности
	isEven := true
	isOdd := true

	// Цикл идёт по steps + 1 точкам от -a до a включительно, с шагом stepSize
	for iter := 0; iter <= steps; iter++ {

		// Вычисляем значение x
		x := -a + stepSize*float64(iter)

		// Вычисляется значение функции f(x) и f(-x) для проверки чётности и нечётности
		fx := givenFunction(x)
		fxNegative := givenFunction(-x)

		// Проверяет, является ли fx значением NaN (Not a Number).
		// NaN может получиться, например, при попытке вычислить math.Sqrt(-1).
		// А еще проверка на бексонечность: math.IsInf(fx, 0) Проверяет, является
		// ли fx бесконечностью (+Inf или -Inf). Например, 1/0.0 даст +Inf.
		if math.IsNaN(fx) || math.IsNaN(fxNegative) || math.IsInf(fx, 0) ||
			math.IsInf(fxNegative, 0) {
			// Если хотя бы одно условие истинно, то печатает следующее сообщение:
			fmt.Printf("Пропуск точки x = %.5f: функция не определена\n", x)
			continue // Не будет проверки на чётность/нечётность для этой точки
		}

		// Проверяем четная функция или нечетная в данной точке:
		if !almostEqual(fx, fxNegative, epsilon) {
			isEven = false
		}
		if !almostEqual(fx, -fxNegative, epsilon) {
			isOdd = false
		}
	}

	fmt.Println("\nРезультат:")
	if isEven {
		fmt.Println("Функция чётная.")
	} else if isOdd {
		fmt.Println("Функция нечётная.")
	} else {
		fmt.Println("Функция ни чётная, ни нечётная.")
	}

	fmt.Print("\n\n")
}
