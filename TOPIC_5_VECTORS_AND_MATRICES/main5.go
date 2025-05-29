package TOPIC_5_VECTORS_AND_MATRICES

import "fmt"

// 1. Строки матрицы A(m,n) заполнены не полностью: в массиве L(m) указано количество
// элементов в каждой строке. Переслать элементы матрицы построчно в начало
// одномерного массива T(m*n) и подсчитать их количество.

// Функция для ввода матрицы и массива L
func inputMatrix(m, n int) ([][]int, []int) {
	//создаем нашу матрицу А и массив с кол-вом элементов в каждой строке
	A := make([][]int, m)
	L := make([]int, m)

	for ii := 0; ii < m; ii++ {
		A[ii] = make([]int, n)

		fmt.Printf("Введите количество заполненных элементов в строке %d (не больше %d): ",
			ii+1, n)
		fmt.Scan(&L[ii])

		if L[ii] > n || L[ii] < 0 {
			fmt.Println("Ошибка: количество заполненных элементов выходит за пределы допустимого.")
			return nil, nil
		}

		for jj := 0; jj < L[ii]; jj++ {
			fmt.Printf("A[%d][%d] = ", ii, jj)
			fmt.Scan(&A[ii][jj])
		}
	}

	return A, L // возвращаем заполненную матрицу А
}

// Функция для пересылки элементов из матрицы A в одномерный массив T
func passMatrix(A [][]int, L []int, m, n int) ([]int, int) {
	T := make([]int, m*n)
	elementCounter := 0

	for ii := 0; ii < m; ii++ {
		for jj := 0; jj < L[ii]; jj++ {
			T[elementCounter] = A[ii][jj]
			elementCounter++
		}
	}
	return T, elementCounter
}

func Main5() {
	fmt.Println("TOPIC_5_VECTORS_AND_MATRICES")
	fmt.Println("Task #1")

	var m, n int
	fmt.Print("Введите количество строк (m): ")
	fmt.Scan(&m)
	fmt.Print("Введите количество столбцов (n): ")
	fmt.Scan(&n)

	// Получаем матрицу A и массив L
	A, L := inputMatrix(m, n)
	if A == nil || L == nil {
		return // выходим из программы
	}

	// Пересылаем данные в одномерный массив
	T, elementCounter := passMatrix(A, L, m, n)

	// Выводим результат
	fmt.Println("\nПересланные элементы матрицы \"А\" в одномерный массив \"T(m*n)\":")
	for iter := 0; iter < elementCounter; iter++ {
		fmt.Printf("%d ", T[iter])
	}

	// Выводим кол-во всех элементов массива T
	fmt.Printf("\nОбщее количество перенесённых элементов: %d\n", elementCounter)

	fmt.Print("\n\n")
}
