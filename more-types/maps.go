package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var mi map[int]Vertex

func main() {
	m = make(map[string]Vertex)
	mi = make(map[int]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	mi[10] = m["Bell Labs"]
	fmt.Println(m["Bell Labs"])
	fmt.Println(mi[10])
}
