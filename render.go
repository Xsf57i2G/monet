package monet

import (
	"image/color"
	"math"

	"github.com/Xsf57i2G/cam"
)

type Renderer struct {
	*buf
	*cam.Camera
}

func (r *Renderer) Render(meshes []*Mesh) []byte {
	r.Clear()
	for _, m := range meshes {
		if m == nil {
			continue
		}
		var radius float64
		for _, v := range m.Vertices {
			var dist = v.position.Length()
			if dist > radius {
				radius = dist
			}
		}
		var w = float64(r.Rect.Max.X)
		var h = float64(r.Rect.Max.Y)
		for i := 0; i < len(m.Indices); i += 3 {
			var i0, i1, i2 = m.Indices[i], m.Indices[i+1], m.Indices[i+2]
			var v0, v1, v2 = m.Vertices[i0], m.Vertices[i1], m.Vertices[i2]
			var s0 = r.Project(v0.position)
			var s1 = r.Project(v1.position)
			var s2 = r.Project(v2.position)
			var x0 = int(s0.X * w)
			var y0 = int(s0.Y * h)
			var x1 = int(s1.X * w)
			var y1 = int(s1.Y * h)
			var x2 = int(s2.X * w)
			var y2 = int(s2.Y * h)
			var dx12 = x1 - x2
			var dy12 = y1 - y2
			var dx20 = x2 - x0
			var dy20 = y2 - y0
			var dx01 = x0 - x1
			var dy01 = y0 - y1
			var minX = int(math.Max(0, math.Min(math.Min(float64(x0), float64(x1)), float64(x2))))
			var maxX = int(math.Min(float64(r.Rect.Max.X-1), math.Max(math.Max(float64(x0), float64(x1)), float64(x2))))
			var minY = int(math.Max(0, math.Min(math.Min(float64(y0), float64(y1)), float64(y2))))
			var maxY = int(math.Min(float64(r.Rect.Max.Y-1), math.Max(math.Max(float64(y0), float64(y1)), float64(y2))))
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					var e1 = dx12*(y-y1) - dy12*(x-x1)
					var e2 = dx20*(y-y2) - dy20*(x-x2)
					var e3 = dx01*(y-y0) - dy01*(x-x0)
					if e1 >= 0 && e2 >= 0 && e3 >= 0 {
						var area = float64(e1 + e2 + e3)
						var a0 = float64(e2) / area
						var a1 = float64(e3) / area
						var a2 = float64(e1) / area
						var color = color.RGBA{
							R: uint8(a0*v0.color.X + a1*v1.color.X + a2*v2.color.X),
							G: uint8(a0*v0.color.Y + a1*v1.color.Y + a2*v2.color.Y),
							B: uint8(a0*v0.color.Z + a1*v1.color.Z + a2*v2.color.Z),
							A: 255,
						}
						r.Set(x, y, color)
					}
				}
			}
		}
	}
	return r.Pix
}
