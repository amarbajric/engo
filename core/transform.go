package core

type Transform struct {
	translation Vector3f
	rotation    Vector3f
	scale       Vector3f
	zNear       float32
	zFar        float32
	width       float32
	height      float32
	fov         float32
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

func (t *Transform) setScale1f(x float64, y float64, z float64) {
	t.scale.X = x
	t.scale.Y = y
	t.scale.Z = z
}

func (t *Transform) setScale3f(vector Vector3f) {
	t.scale = vector
}

func (t *Transform) setProjection(fov float64, width float64, height float64, zNear float64, zFar float64) {
	t.fov = float32(fov)
	t.width = float32(width)
	t.height = float32(height)
	t.zNear = float32(zNear)
	t.zFar = float32(zFar)
}

func (t *Transform) getTransformation() Matrix4f {
	translationMatrix := Matrix4f{[4][4]float64{}}
	rotationMatrix := Matrix4f{[4][4]float64{}}
	scaleMatrix := Matrix4f{[4][4]float64{}}
	translationMatrix.InitTranslation(t.translation.X, t.translation.Y, t.translation.Z)
	rotationMatrix.InitRotation(t.rotation.X, t.rotation.Y, t.rotation.Z)
	scaleMatrix.InitScale(t.scale.X, t.scale.Y, t.scale.Z)

	scaledRotationMatrix := rotationMatrix.Mul(&scaleMatrix)
	return translationMatrix.Mul(&scaledRotationMatrix)
}

func (t *Transform) getProjectedTransformation() Matrix4f {
	transformationMatrix := t.getTransformation()
	projectionMatrix := Matrix4f{M: [4][4]float64{}}
	projectionMatrix.InitProjection(float64(t.fov), float64(t.width), float64(t.height), float64(t.zNear), float64(t.zFar))

	return projectionMatrix.Mul(&transformationMatrix)
}
