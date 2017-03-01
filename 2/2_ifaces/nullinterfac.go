package main

import "fmt"

func main() {
	var f interface{}

	fmt.Printf("%v %v", f, f == nil)

	if f = getNil(10); f == nil {
		fmt.Println("i'm nill")
	}
}

func getNil(input interface{}) interface{} {
	if _, ok := input.(int); ok {
		return struct{}{}
	}
	return nil
}
