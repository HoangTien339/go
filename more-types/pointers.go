package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	var k *int

	p := &i                           // point to i
	fmt.Printf("Address p: %v \n", p) // print value of pointer
	fmt.Println(*p)                   // read i through the pointer

	k = new(int) // point to p
	*k = 22
	fmt.Printf("Address k: %v \n", *k) // print value of pointer

	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
