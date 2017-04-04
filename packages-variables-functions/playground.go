package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(100))

	fmt.Println("Now you have %g proplems.", math.Sqrt(7))

	fmt.Println(math.Pi)

}
