package main

// Первый закон рефлексии, на входе всегда идет Interface{}

import (
	"fmt"
	"reflect"
)

func main() {
	x := 3.4
	// TypeOf() принимает на вход interface{}, в этом месте будет аллокация
	fmt.Println("reflect.Type:", reflect.TypeOf(x))

	// reflect.Value != значению переданному на вход
	fmt.Println("reflect.Value:", reflect.ValueOf(x).String())

	v := reflect.ValueOf(x)
	fmt.Println("Тип value:", v.Type())
	fmt.Println("тип float64:", v.Kind() == reflect.Float64)
	fmt.Println("Значение:", v.Float())

	// Важно отметить, что Kind - это базовый тип (int/float/struct/slice)
	// А не пользовательский тип
	type MyInt int
	var c MyInt = 7
	v = reflect.ValueOf(c)
	fmt.Println("kind is int: ", v.Kind() == reflect.Int) // true.
}
