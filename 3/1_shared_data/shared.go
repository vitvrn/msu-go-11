package main

import (
	"fmt"
	"math/rand"
)

func fillRandom(box []int, i int) {
	// рандомный индекс [0,i)
	box[rand.Intn(i)]++
}

func main() {
	rand.Seed(429)
	stuff := make([]int, 10)

	for i := 0; i < 10; i++ {
		go fillRandom(stuff, 10)
	}
	fmt.Scanln()
	fmt.Println(stuff)

	// closure()
}

func closure() {
	// Функции захватывают переменные в области видимости
	// Но, чтобы передавать значение, требуется явно передавать его в функцию
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("Got", i)
		}()
	}
	fmt.Scanln()
}
