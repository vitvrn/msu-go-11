package main

import "fmt"

var c chan int

func main() {
	c := make(chan string)
	go greet(c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c, ",", <-c)
	}
}

func greet(c chan string) {
	for {
		c <- fmt.Sprintf("Владыка")
		c <- fmt.Sprintf("Штурмовик")
	}
}
