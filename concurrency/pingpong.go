package main

import (
	"fmt"
	"time"
)

type Ball struct{ hits int }

func main() { // main goroutine
	table := make(chan *Ball)
	go player("ping", table) // goroutine2
	go player("pong", table) // goroutine3
	go player("pang", table) // goroutine4

	table <- new(Ball) // game on; toss the ball
	time.Sleep(2 * time.Second)
	<-table // game over; grab the ball
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
