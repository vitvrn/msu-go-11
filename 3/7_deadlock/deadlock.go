package main

import (
	"fmt"
	"time"
)

// Ball is just a ball
type Ball struct{ hits int }

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	// table <- new(Ball) // Запуска мяча в игру
	time.Sleep(1 * time.Second)
	<-table // Конец игры, забираем мяч
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
