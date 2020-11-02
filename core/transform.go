package core

type Transform struct {
	translation Vector3f
	rotation    Vector3f
}

func (t *Transform) setTranslation1f(x float64, y float64, z float64) {
	t.translation.X = x
	t.translation.Y = y
	t.translation.Z = z
}

func (t *Transform) setTranslation3f(vector Vector3f) {
	t.translation = vector
}

func (t *Transform) setRotation1f(x float64, y float64, z float64) {
	t.rotation.X = x
	t.rotation.Y = y
	t.rotation.Z = z
}

func (t *Transform) setRotation3f(vector Vector3f) {
	t.rotation = vector
}

func (t *Transform) getTransformation() Matrix4f {
	translationMatrix := Matrix4f{M: [4][4]float64{}}
	rotationMatrix := Matrix4f{M: [4][4]float64{}}
	translationMatrix.InitTranslation(t.translation.X, t.translation.Y, t.translation.Z)
	rotationMatrix.InitRotation(t.rotation.X, t.rotation.Y, t.rotation.Z)
	return translationMatrix.Mul(&rotationMatrix)
}
