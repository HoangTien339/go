package main

import (
	"fmt"
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (m MyReader) Read(b []byte) (int, error) {
	fmt.Println(len(b))
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
