package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	bounds image.Rectangle
	data   [][]color.RGBA
}

func (i Image) Bounds() image.Rectangle {
	return i.bounds
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return i.data[y][x]
}

func NewImage(dx, dy int) *Image {
	p := make([][]color.RGBA, dy)
	for i := range p {
		p[i] = make([]color.RGBA, dx)
	}

	xStep := float32(255.0) / float32(dx+dy-1)

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			intensity := uint8(float32(x+y) * xStep)
			c := color.RGBA{intensity, intensity, intensity, 255}
			p[y][x] = c
		}
	}

	return &Image{image.Rectangle{image.Point{0, 0}, image.Point{dx, dy}}, p}
}

func main() {
	m := NewImage(100, 100)
	pic.ShowImage(m)
}
