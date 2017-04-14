package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
	c    uint8
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) At(x, y int) color.Color {
	return color.RGBA{i.c + uint8(x), i.c + uint8(y), 255, 255}
}

func main() {
	m := Image{100, 100, 255}
	pic.ShowImage(&m)
}
