package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func run() error {
	var r *MyError = &MyError{
		time.Now(),
		"it didnt work!",
	}
	return r
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
