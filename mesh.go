package monet

import (
	"github.com/Xsf57i2G/geom"
)

type vertex struct {
	position, normal, color geom.Vec
}

type Mesh struct {
	Vertices []vertex
	Indices  []uint32
	Normals  []geom.Vec
	Color    []geom.Vec
}

func (m *Mesh) Shade() []float64 {
	return nil
}
