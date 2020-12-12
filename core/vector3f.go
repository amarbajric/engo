package core

import (
	"fmt"
	"math"
)

type Vector3f struct {
	X float64
	Y float64
	Z float64
}

func (v *Vector3f) ToString() string {
	return fmt.Sprintf("(%.2f, %.2f, %.2f)", v.X, v.Y, v.Z)
}

func (v *Vector3f) Length() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

func (v *Vector3f) Dot(r *Vector3f) float64 {
	return (v.X * r.X) + (v.Y * r.Y) + (v.Z * r.Z)
}

func (v *Vector3f) Cross(r *Vector3f) Vector3f {
	_x := v.Y * r.Z - v.Z * r.Y
	_y := v.Z * r.X - v.X * r.Z
	_z := v.X * r.Y - v.Y * r.X

	return Vector3f{
		X: _x,
		Y: _y,
		Z: _z,
	}
}

func (v *Vector3f) Normalize() {
	length := v.Length()
	v.X /= length
	v.Y /= length
	v.Z /= length
}

func (v *Vector3f) Rotate(angle float64, axis Vector3f) {
	sinHalfAngle := math.Sin((angle / 2) * (math.Pi/180))
	cosHalfAngle := math.Cos((angle / 2) * (math.Pi/180))

	rX := axis.X * sinHalfAngle
	rY := axis.Y * sinHalfAngle
	rZ := axis.Y * sinHalfAngle
	rW := cosHalfAngle

	rotation := Quaternion{
		X: rX,
		Y: rY,
		Z: rZ,
		W: rW,
	}
	conjugate := rotation.Conjugate()

	w := rotation.MulV(*v)
	w2 := w.Mul(conjugate)

	v.X = w2.X
	v.Y = w2.Y
	v.Z = w2.Z
}

func (v *Vector3f) AddV(r Vector3f) Vector3f {
	return Vector3f{
		X: v.X + r.X,
		Y: v.Y + r.Y,
		Z: v.Z + r.Z,
	}
}

func (v *Vector3f) AddF(r float64) Vector3f {
	return Vector3f{
		X: v.X + r,
		Y: v.Y + r,
		Z: v.Z + r,
	}
}

func (v *Vector3f) SubV(r Vector3f) Vector3f {
	return Vector3f{
		X: v.X - r.X,
		Y: v.Y - r.Y,
		Z: v.Z - r.Z,
	}
}

func (v *Vector3f) SubF(r float64) Vector3f {
	return Vector3f{
		X: v.X - r,
		Y: v.Y - r,
		Z: v.Z - r,
	}
}

func (v *Vector3f) MulV(r Vector3f) Vector3f {
	return Vector3f{
		X: v.X * r.X,
		Y: v.Y * r.Y,
		Z: v.Z * r.Z,
	}
}

func (v *Vector3f) MulF(r float64) Vector3f {
	return Vector3f{
		X: v.X * r,
		Y: v.Y * r,
		Z: v.Z * r,
	}
}

func (v *Vector3f) DivV(r Vector3f) Vector3f {
	return Vector3f{
		X: v.X / r.X,
		Y: v.Y / r.Y,
		Z: v.Z / r.Z,
	}
}

func (v *Vector3f) DivF(r float64) Vector3f {
	return Vector3f{
		X: v.X / r,
		Y: v.Y / r,
		Z: v.Z / r,
	}
}
