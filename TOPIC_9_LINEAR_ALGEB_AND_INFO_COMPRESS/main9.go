package TOPIC_9_LINEAR_ALGEB_AND_INFO_COMPRESS

import "fmt"

// 1. Выполнить операцию транспонирования прямоугольной матрицы A(m,n), m!=n,
// не выделяя дополнительного массива для хранения результата.
// Матрицу представить в виде одномерного массива.

// Выводим транспонированную матрицу размером rows*cols, представленную одномерным массивом
func printMatrixTranspose(matrix []int, rows, cols int) {
	for ii := 0; ii < cols; ii++ {
		for jj := 0; jj < rows; jj++ {
			fmt.Printf("%d ", matrix[jj*cols+ii])
		}
		fmt.Println()
	}

}

// Выводим начальную матрицу размером rows*cols, представленную одномерным массивом
func printMatrix(matrix []int, rows, cols int) {
	for ii := 0; ii < rows; ii++ {
		for jj := 0; jj < cols; jj++ {
			fmt.Printf("%d ", matrix[ii*cols+jj])
		}
		fmt.Println()
	}
}

func Main9() {
	fmt.Println("TOPIC_9_LINEAR_ALGEB_AND_INFO_COMPRESS")
	fmt.Println("Task #1")

	var m, n int
	fmt.Print("Введите число строк m: ")
	fmt.Scan(&m)
	fmt.Print("Введите число столбцов n: ")
	fmt.Scan(&n)

	if m == n {
		fmt.Println("Матрица квадратная, что противоречит условию задачи m != n")
		return
	}

	length := m * n
	A := make([]int, length)

	fmt.Printf("Введите элементы матрицы A(%d x %d) построчно:\n", m, n)
	for iter := 0; iter < length; iter++ {
		fmt.Scan(&A[iter])
	}

	fmt.Printf("Исходная матрицаA(%d x %d) построчно:\n", m, n)
	printMatrix(A, m, n)

	fmt.Printf("Транспонированная матрица (%d x %d):\n", n, m)
	printMatrixTranspose(A, m, n)

	fmt.Printf("\n\n")
}
