package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width  int
	height int
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.width, m.height)
}

func (m Image) At(x, y int) color.Color {
	v := uint8((x^y + x^y)/2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{1024, 512}
	pic.ShowImage(m)
}
