package core

import (
	"fmt"
	"math"
)

type Vector2f struct {
	X float64
	Y float64
}

func (v *Vector2f) ToString() string {
	return fmt.Sprintf("(%.2f, %.2f)", v.X, v.Y)
}

func (v *Vector2f) Length() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v *Vector2f) Dot(r *Vector2f) float64 {
	return (v.X * r.X) + (v.Y * r.Y)
}

func (v *Vector2f) Normalize() {
	length := v.Length()
	v.X /= length
	v.Y /= length
}

func (v *Vector2f) Rotate(angle float64) Vector2f {
	rad := angle * (math.Pi / 180)
	cos := math.Cos(rad)
	sin := math.Sin(rad)

	return Vector2f{
		X: v.X * cos - v.Y * sin,
		Y: v.X * sin + v.Y * cos,
	}
}

func (v *Vector2f) AddV(r Vector2f) Vector2f {
	return Vector2f{
		X: v.X + r.X,
		Y: v.Y + r.Y,
	}
}

func (v *Vector2f) AddF(r float64) Vector2f {
	return Vector2f{
		X: v.X + r,
		Y: v.Y + r,
	}
}

func (v *Vector2f) SubV(r Vector2f) Vector2f {
	return Vector2f{
		X: v.X - r.X,
		Y: v.Y - r.Y,
	}
}

func (v *Vector2f) SubF(r float64) Vector2f {
	return Vector2f{
		X: v.X - r,
		Y: v.Y - r,
	}
}

func (v *Vector2f) MulV(r Vector2f) Vector2f {
	return Vector2f{
		X: v.X * r.X,
		Y: v.Y * r.Y,
	}
}

func (v *Vector2f) MulF(r float64) Vector2f {
	return Vector2f{
		X: v.X * r,
		Y: v.Y * r,
	}
}

func (v *Vector2f) DivV(r Vector2f) Vector2f {
	return Vector2f{
		X: v.X / r.X,
		Y: v.Y / r.Y,
	}
}

func (v *Vector2f) DivF(r float64) Vector2f {
	return Vector2f{
		X: v.X / r,
		Y: v.Y / r,
	}
}
