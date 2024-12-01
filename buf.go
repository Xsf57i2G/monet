package monet

import (
	"image"
	"image/color"
)

type buf struct {
	image.RGBA
	z   []float64
	fog []float64
}

func (b *buf) Clear() {
	for i := range b.Pix {
		b.Pix[i] = 0
	}
	for i := range b.z {
		b.z[i] = 1
	}
}

func (b *buf) Draw(x, y int, z float64, c color.RGBA) {
	var index = y*b.Rect.Max.X + x
	if z < b.z[index] {
		b.z[index] = z
		b.Set(x, y, c)
	}
}
