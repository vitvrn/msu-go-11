package main

import (
	"fmt"
	"reflect"
)

// Второй закон рефлексии
// Чтобы работать с подлежащим значением, нужно получить interface{} обратно

func main() {
	type MyT float64
	var x MyT = 3.4
	v := reflect.ValueOf(x)

	y := v.Interface().(MyT) // y will have type float64.
	fmt.Println("Значение обертки", v, "Само значение", y)
}
