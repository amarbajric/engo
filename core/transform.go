package core

type Transform struct {
	translation Vector3f
}

func (t *Transform) setTranslation1f(x float64, y float64, z float64) {
	t.translation.X = x
	t.translation.Y = y
	t.translation.Z = z
}

func (t *Transform) setTranslation3f(vector Vector3f) {
	t.translation.X = vector.X
	t.translation.Y = vector.Y
	t.translation.Z = vector.Z
}

func (t *Transform) getTransformation() Matrix4f {
	translationMatrix := Matrix4f{M: [4][4]float64{}}
	translationMatrix.InitTranslation(t.translation.X, t.translation.Y, t.translation.Z)
	return translationMatrix
}
