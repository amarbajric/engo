package core

const VERTEX_SIZE = 3

type Vertex struct {
	Pos  Vector3f
}

func (v *Vertex) Init(vt Vector3f) {
	v.Pos = vt
}
