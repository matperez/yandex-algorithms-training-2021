package main

import (
	"fmt"
)

func main() {
	var a1, b1, a2, b2 int
	fmt.Scan(&a1, &b1, &a2, &b2)
	x, y := solve(a1, b1, a2, b2)
	fmt.Print(x, y)
}

//solve решает задачу
func solve(a1, b1, a2, b2 int) (int, int) {
	var (
		items [][]int
		minAreaItem []int
		minArea int
	)
	// поворачиваем прямоугольники, чтобы оба стояли
	a1, b1 = rotate(a1, b1)
	a2, b2 = rotate(a2, b2)
	// накидываем варианты компоновки
	items = append(items, []int{max(a1, a2), b1 + b2})
	items = append(items, []int{max(b1, b2), a1 + a2})
	items = append(items, []int{max(b1, a2), a1 + b2})
	items = append(items, []int{max(b2, a1), a2 + b1})
	// вычисляем вариант с минимальной площадью
	minArea = 2000 * 2000
	for _, item := range items {
		itemArea := item[0] * item[1]
		if  itemArea <= minArea {
			minAreaItem = item
			minArea = itemArea
		}
	}
	return minAreaItem[0], minAreaItem[1]
}

//max максимум из двух чисел
func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

//rotate поворачивает прямоугольник в стоячее положение
func rotate(a, b int) (int, int) {
	if a >= b {
		return a, b
	} else {
		return b, a
	}
}
