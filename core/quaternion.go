package core

import "math"

type Quaternion struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (q *Quaternion) Length() float64 {
	return math.Sqrt(math.Pow(q.X, 2) + math.Pow(q.Y, 2) + math.Pow(q.Z, 2) + math.Pow(q.W, 2))
}

func (q *Quaternion) Normalize() {
	length := q.Length()

	q.X /= length
	q.Y /= length
	q.Z /= length
	q.W /= length
}

func (q *Quaternion) Conjugate() Quaternion {
	return Quaternion{
		X: -q.X,
		Y: -q.Y,
		Z: -q.Z,
		W: -q.W,
	}
}

func (q *Quaternion) Mul(r Quaternion) Quaternion {
	_x := q.X * r.W + q.W * r.X + q.Y * r.Z - q.Z * r.Y
	_y := q.Y * r.W + q.W * r.Y + q.Z * r.X - q.X * r.Z
	_z := q.Z * r.W + q.W * r.Z + q.X * r.Y - q.Y * r.X
	_w := q.W * r.W - q.X * r.X - q.Y * r.Y - q.Z * r.Z

	return Quaternion{
		X: _x,
		Y: _y,
		Z: _z,
		W: _w,
	}
}

func (q *Quaternion) MulV(r Vector3f) Quaternion {
	_x := q.W * r.X + q.Y * r.Z - q.Z * r.Y
	_y := q.W * r.Y + q.Z * r.X - q.X * r.Z
	_z := q.W * r.Z + q.X * r.Y - q.Y * r.X
	_w := -q.X * r.X - q.Y * r.Y - q.Z * r.Z

	return Quaternion{
		X: _x,
		Y: _y,
		Z: _z,
		W: _w,
	}
}