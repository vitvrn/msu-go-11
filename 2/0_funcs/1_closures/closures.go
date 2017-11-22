package main

import (
	"fmt"
	"time"
)

func main() {
	myTimer := getTimer()
	//vit: на лекции удалена строка:
	defer myTimer()

	//vit:
	greet := "hello from closure #2?"

	f := func() {
		println(greet)
		//TODO ^^^ ??? почему-то выводится то до, то после "Time from start..." !!!
		//!!! В других программах то же (распараллеливание?)
		myTimer()
	}

	//vit (если сделаем так, то f не вызывается => не выводится "Time from start..."):
	//	f()
	//println(f)

	f()
}

func getTimer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time from start %v\n", time.Since(start))
	}
}
