package main

import (
	"fmt"
)

func main() {
	var n, k, m int
	fmt.Scan(&n, &k, &m)
	x := solve(n, k, m)
	fmt.Print(x)
}

//solve решает задачу
func solve(n, k, m int) int {
	var x, d, r int
	for {
		d, r = produce(n, k, m)
		x += d
		n = r
		if d == 0 {
			break
		}
	}
	return x
}

//produce производит работу
func produce(n, k, m int) (int, int) {
	// число заготовок из имеющегося материала
	kn := n / k
	if kn < 1 {
		return 0, 0
	}
	// остаток после изготовления заготовок
	kr := n % k
	mn := k / m * kn
	mr := k % m * kn
	return mn, mr + kr
}
