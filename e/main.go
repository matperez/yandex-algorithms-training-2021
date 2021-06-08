package main

import (
	"fmt"
	"math"
)

func main() {
	var k1, k2, m, p1, p2, n1, n2 int64
	fmt.Scan(&k1, &m, &k2, &p2, &n2)
	p1, n1 = solve(k1, m, k2, p2, n2)
	fmt.Print(p1, n1)
}

// k1, k2 - номера квартир
// m - число этажей в доме
// p1, p2 - номера подъездов
// n1, n2 - номера этажей
func solve(k2, m, k1, p1, n1 int64) (int64, int64) {
	var n2, p2 int64

	// если данные противоречивы
	if !valid(m, k1, p1, n1) {
		return -1, -1
	}

	// число квартир на этаж
	apf := int64(math.Ceil(float64(k1) / float64(m * (p1 - 1) + n1)))

	// подъезд
	p2 = int64(math.Ceil(float64(k2) / float64(m * apf)))

	// этаж
	n2 = (k2 - ((p2 - 1) * m * apf)) / apf + 1

	fmt.Println(k2, m * (p2 - 1) * apf + (n2 - 1) * apf, k2 >= m * (p2 - 1) * apf + (n2 - 1) * apf)

	//if n2 > m {
	//	p2 = 0
	//	n2 = m
	//}

	return p2, n2
}

//valid проверка валидности исходных данных
// номер известной квартиры связан с количеством квартир на этаже следующим соотношением
// k >= m * (p - 1) * x + n * x, где x - количество квартир на этаже
// в то же время количество квартир на этаже не может быть меньше 1
// получаем неравенство 1 <= x <= k / (m * (p - 1) + n)
// из которого, сократив x можно получить, что k / (m * (p - 1) + n) >= 1
func valid(m, k, p, n int64) bool {
	return k >= m * (p - 1) + n
}
