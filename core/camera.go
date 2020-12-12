package core

type Camera struct {
	pos     Vector3f
	forward Vector3f
	up      Vector3f
	yAxis   Vector3f
}

func (cam *Camera) InitCamera() {
	cam.pos = Vector3f{0, 0, 0}
	cam.forward = Vector3f{0, 0, 1}
	cam.up = Vector3f{0, 1, 0}
	cam.yAxis = Vector3f{0, 1, 0}
}

func (cam *Camera) Move(dir Vector3f, amount float64) {
	cam.pos = cam.pos.AddV(dir.MulF(amount))
}

func (cam *Camera) RotateX(angle float64) {
	hAxis := cam.yAxis.Cross(&cam.forward)
	hAxis.Normalize()

	cam.forward.Rotate(angle, hAxis)
	cam.forward.Normalize()

	cam.up = cam.forward.Cross(&hAxis)
	cam.up.Normalize()
}

func (cam *Camera) RotateY(angle float64) {
	hAxis := cam.yAxis.Cross(&cam.forward)
	hAxis.Normalize()

	cam.forward.Rotate(angle, cam.yAxis)
	cam.forward.Normalize()

	cam.up = cam.forward.Cross(&hAxis)
	cam.up.Normalize()
}

func (cam *Camera) GetLeft() Vector3f {
	left := cam.forward.Cross(&cam.up)
	left.Normalize()
	return left
}

func (cam *Camera) GetRight() Vector3f {
	right := cam.up.Cross(&cam.forward)
	right.Normalize()
	return right
}