package monet

import (
	"image/color"

	"github.com/Xsf57i2G/geom"
)

type Light struct {
	Position    geom.Vec
	Color       color.Color
	Intensity   float64
	Attenuation float64
}

func (l *Light) Illuminate(p geom.Vec) (color.Color, float64) {
	var d = p.Sub(l.Position)
	var dist = d.Length()
	var intensity = l.Intensity / (1 + l.Attenuation*dist*dist)
	return l.Color, intensity
}
