package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %v\n", b[:n]) // fmt.Printf("b[:n] = %v\n", string(b[:n]))
		if err == io.EOF {
			fmt.Println("-", io.EOF, "-")
			break
		}
	}
}
