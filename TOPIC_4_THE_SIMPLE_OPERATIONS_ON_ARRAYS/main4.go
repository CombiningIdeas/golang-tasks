package TOPIC_4_THE_SIMPLE_OPERATIONS_ON_ARRAYS

import (
	"fmt"
)

// 1. В массиве C(n) подсчитать количество отрицательных и сумму положительных элементов.

func Main4() {
	fmt.Println("TOPIC_4_THE_SIMPLE_OPERATIONS_ON_ARRAYS")
	fmt.Println("Task #1")

	var lenArray int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scan(&lenArray)

	fmt.Println("Введите элементы массива:")

	//создание массива, в котором будем все считать
	soutceArray := make([]int, lenArray)

	for iter := 0; iter < lenArray; iter++ {
		fmt.Printf("Элемент %d: ", iter+1)
		fmt.Scan(&soutceArray[iter])
	}

	var countNegatives int // Количество отрицательных
	var sumPositives int   // Сумма положительных

	// Проходимся по массиву:
	for _, value := range soutceArray {
		if value < 0 {
			countNegatives++
		} else if value > 0 {
			sumPositives += value
		}
	}

	// Выводим результат:
	fmt.Printf("Количество отрицательных элементов: %d\n", countNegatives)
	fmt.Printf("Сумма положительных элементов: %d\n", sumPositives)

	fmt.Print("\n\n")
}
