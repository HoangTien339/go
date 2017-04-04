package main

import (
	"fmt"
)

func print() int {
	fmt.Println("Good food ! Today weather is so nice !")
	return 0
}

func main() {
	var i int = 0
	switch i {
	case 0:
	case print():
		fmt.Println("You are stupid")
	default:
		fmt.Println("You are fool.")
	}
}
