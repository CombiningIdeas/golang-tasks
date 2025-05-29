package TOPIC_8_GEOMETRY_AND_SET_THEORY

import "C"
import (
	"fmt"
	"math"
	"sort"
)

// 1. Заяц, хаотично прыгая, оставил след в виде замкнутой самопересекающейся ломаной,
// охватывающей территорию его владения (отрезки ломаной заданы длиной прыжка и его
// направлением по азимуту). Найти площадь минимального по площади выпуклого
// многоугольника, описанного вокруг этой территории.

// Будем использовать Метод Грэхема (также называемый методом сканирования
// Грэхема или обходом Грэхема) — это алгоритм для построения выпуклой оболочки
// конечного набора точек на плоскости. Его асимптотическая сложность составляет O(n log n),
// где n — число точек.
// Так же будем использовать вычисление площади многоугольника по формуле Гаусса.

// Структура хранящая кординаты на плоскости в декартовой системе
type Point struct {
	x, y float64
}

// Перевод из полярных координат (длина, азимут) в декартовые (x, y)
func polarToCartesian(start Point, length, azimuth float64) Point {
	radians := azimuth * math.Pi / 180

	// Записываем результат перевода в структуру,
	// каждый раз при вызове этой функции в новую структуру
	return Point{
		x: start.x + length*math.Cos(radians),
		y: start.y + length*math.Sin(radians),
	}
}

// Вычисление векторного произведения для проверки ориентации трех точек
func cross(a, b, c Point) float64 {
	// Это значение показывает, с какой стороны от вектора AB лежит точка C.
	return (b.x-a.x)*(c.y-a.y) - (b.y-a.y)*(c.x-a.x)
	// > 0 — точка c лежит слева от направления из a в b (против часовой стрелки)
	// < 0 — точка c лежит справа (по часовой стрелке)
	// = 0 — точки коллинеарны (на одной прямой).
}

// Построение выпуклой оболочки методом Грэхема
func convexHull(points []Point) []Point {
	// Сортируем точки по x, при равенстве — по y
	sort.Slice(points, func(i, j int) bool {
		// Используем встроенную сортировку на срезе points,
		// массиве структур Point (точек с координатами x и y).
		if points[i].x == points[j].x {
			return points[i].y < points[j].y
		}
		return points[i].x < points[j].x
	})

	// lower - Нижняя оболочка, часть многоугольника, которая образует «нижнюю» сторону
	//оболочки, проходящую слева направо по нижней границе множества точек.
	// upper - Верхняя оболочка, часть многоугольника, проходящая слева направо по
	// верхней границе множества точек.
	var lower, upper []Point

	// Строим нижнюю часть оболочки
	for _, p := range points {
		for len(lower) >= 2 && cross(lower[len(lower)-2], lower[len(lower)-1], p) <= 0 {
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, p)
	}

	// Строим верхнюю часть оболочки
	for ii := len(points) - 1; ii >= 0; ii-- {
		p := points[ii]
		for len(upper) >= 2 && cross(upper[len(upper)-2], upper[len(upper)-1], p) <= 0 {
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, p)
	}

	// Объединяем части, убирая дублирующиеся концы
	return append(lower[:len(lower)-1], upper[:len(upper)-1]...)
}

// Вычисление площади многоугольника по формуле Гаусса
func polygonArea(polygon []Point) float64 {
	area := 0.0
	countPoints := len(polygon) // Количество точек в многоугольнике
	for ii := 0; ii < countPoints; ii++ {
		jj := (ii + 1) % countPoints // Индекс следующей точки (с циклическим переходом на первую)
		area += polygon[ii].x*polygon[jj].y - polygon[jj].x*polygon[ii].y
	}
	return math.Abs(area) / 2
}

func Main8() {
	fmt.Println("TOPIC_8_GEOMETRY_AND_SET_THEORY")
	fmt.Println("Task #1")

	var numberOfJumps int // Кол-во прыжков зайца
	fmt.Print("Введите количество прыжков: ")
	fmt.Scan(&numberOfJumps)

	var points []Point
	current := Point{0, 0}
	points = append(points, current)

	fmt.Println("Введите длину и азимут каждого прыжка через пробел:")
	// Формируем последовательность точек по входным данным
	for iter := 0; iter < numberOfJumps; iter++ {
		var length, azimuth float64
		fmt.Scan(&length, &azimuth)
		current = polarToCartesian(current, length, azimuth)
		points = append(points, current)
	}

	shellFigure := convexHull(points)      // Строим выпуклую оболочку по полученным точкам
	areaFigure := polygonArea(shellFigure) // Вычисляем площадь выпуклого многоугольника

	fmt.Printf("Площадь выпуклой оболочки: %.8f\n", areaFigure)

	fmt.Printf("\n\n")
}
