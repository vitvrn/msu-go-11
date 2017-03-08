package main

import (
	"fmt"
)

func main() {
	fmt.Println("старт")
	// можно запустить функцию
	go process(0)
	// можно запустить анонимную функцию
	go func() {
		fmt.Println("Анонимный запуск")
	}()

	// Можем запустить много горутин
	for i := 0; i < 1000; i++ {
		go process(i)
	}

	// ждем
	var input string
	fmt.Scanln(&input)

}

func process(i int) {
	fmt.Println("обработка: ", i)
}
