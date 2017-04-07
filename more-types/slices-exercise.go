package main

import (
	"golang.org/x/tour/pic"
	"math/rand"
)

const MaxUint = int(^uint8(0))

func Pic(dx, dy int) [][]uint8 {
	colors := make([][]uint8, dx)
	for y, _ := range colors {
		colors[y] = make([]uint8, dx)
		for x, _ := range colors[y] {
			dot := rand.Intn(MaxUint)
			// Can replace dot = (x + y)/2 | x*y | x^y
			colors[y][x] = uint8(dot)
		}
	}
	return colors
}

func main() {
	pic.Show(Pic)
}
