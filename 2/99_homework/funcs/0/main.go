package main

import (
	"fmt"
)

// TODO: Реализовать вычисление Квадратного корня
func Sqrt(A float64) float64 {
	if A < 0 {
		return 0 //??????? how to return NaN ??????????????
	} else if A == 0 {
		return 0
	}
	epsilon := 0.01
	var x, x0 float64
	x0 = A
	for {
		x = (x0 + A/x0) / 2
		fmt.Println("x=", x, ", x0=", x0)
		if x >= x0 && x-x0 < epsilon || x < x0 && x0-x < epsilon { //можно убрать половину условия если всегда x<=x0?
			break
		}
		x0 = x
	}

	//	println(x0, epsilon)

	return x
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(100))
}
