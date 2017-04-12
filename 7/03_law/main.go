package main

import (
	"fmt"
	"reflect"
)

// Третий закон рефлексии
// Свойство "settable"

func main() {
	var x float64 = 3.4
	// Мы создаем копию
	v := reflect.ValueOf(x)
	// Изменинеи v запрещено, тк отсутствует связь с подлежащим значением
	v.SetFloat(7.1) // Error: will panic.

	fmt.Println("settability of v:", v.CanSet())

	// Чтобы иметь возможность изменить значение, нам потребуется ссылка
	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

	// Теперь, использую Elem мы получим Value, лежащее по ссылке
	v = p.Elem()
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}
