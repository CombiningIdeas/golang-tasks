package TOPIC_6_LINEAR_SEARCH

import "fmt"

// 1. Седловой точкой в матрице называется элемент, являющийся одновременно
// наибольшим в столбце и наименьшим в строке. Седловых точек может быть несколько.
// В матрице A(m,n) найти все седловые точки либо установить, что таких точек нет.

// Возвращает список седловых точек в виде среза [значение, строка, столбец]
func FindSaddlePoints(matrix [][]int) [][3]int {
	rows := len(matrix)
	cols := len(matrix[0])
	var resultArray [][3]int

	for ii := 0; ii < rows; ii++ {
		// Ищем минимум в строке i
		minValue := matrix[ii][0]
		minCol := 0
		for jj := 1; jj < cols; jj++ {
			if matrix[ii][jj] < minValue {
				minValue = matrix[ii][jj]
				minCol = jj
			}
		}

		// Проверим, является ли точка максимумом в своём столбце
		isSaddle := true
		for ii := 0; ii < rows; ii++ {
			if matrix[ii][minCol] > minValue {
				isSaddle = false
				break
			}
		}

		// Если флаг истинен значит седловая точка есть, добавляем её и её позицию в матрице
		if isSaddle {
			// Добавляем один элемент (массив из трёх чисел [3]int),
			// поэтому мы можно сказать не распаковываем его и не пишем в конце ...
			// Так как resultArray двумерный, а не одномерный
			resultArray = append(resultArray, [3]int{minValue, ii, minCol})
		}
	}

	return resultArray
}

func outputArray(matrix [][]int, rows, cols int) {
	for ii := 0; ii < rows; ii++ {
		for jj := 0; jj < rows; jj++ {
			fmt.Printf("%d ", matrix[ii][jj])
		}
		fmt.Println()
	}
}

func Main6() {
	fmt.Println("TOPIC_6_LINEAR_SEARCH")
	fmt.Println("Task #1")

	fmt.Println("TOPIC_5_VECTORS_AND_MATRICES")
	fmt.Println("Task #1")

	var rows, cols int
	fmt.Print("Введите количество строк: ")
	fmt.Scan(&rows)
	fmt.Print("Введите количество столбцов: ")
	fmt.Scan(&cols)

	if rows <= 0 || cols <= 0 {
		fmt.Println("Размеры матрицы должны быть положительными")
		return
	}

	// Ввод матрицы
	matrix := make([][]int, rows)
	for ii := range matrix {
		matrix[ii] = make([]int, cols)
		for jj := 0; jj < cols; jj++ {
			fmt.Printf("Введите элемент [%d][%d]: ", ii, jj)
			fmt.Scan(&matrix[ii][jj])
		}
		fmt.Print("\n")
	}

	// Поиск седловых точек
	saddles := FindSaddlePoints(matrix)

	// Вывод введённой матрицы для наглядности
	outputArray(matrix, rows, cols)

	if len(saddles) == 0 {
		fmt.Println("Седловых точек не найдено.")
	} else {
		for _, s := range saddles {
			fmt.Printf("Седловая точка: %d - на позиции (%d, %d)\n", s[0], s[1], s[2])
		}
	}

	fmt.Print("\n\n")
}
