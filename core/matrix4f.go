package core

import "math"

type Matrix4f struct {
	M [4][4]float64
}

func (m *Matrix4f) InitIdentity() {
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if i == j {
				m.M[i][j] = 1
			} else {
				m.M[i][j] = 0
			}
		}
	}
}

func (m *Matrix4f) InitTranslation(x float64, y float64, z float64) {
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if i == j {
				m.M[i][j] = 1
			} else if j == 3 {
				if i == 0 {
					m.M[i][j] = x
				} else if i == 1 {
					m.M[i][j] = y
				} else if i == 2 {
					m.M[i][j] = z
				}
			} else {
				m.M[i][j] = 0
			}
		}
	}
}

func (m *Matrix4f) InitScale(x float64, y float64, z float64) {
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if i == j {
				if i == 0 {
					m.M[i][j] = x
				} else if i == 1 {
					m.M[i][j] = y
				} else if i == 2 {
					m.M[i][j] = z
				} else {
					m.M[i][j] = 1
				}
			} else {
				m.M[i][j] = 0
			}
		}
	}
}

func (m *Matrix4f) InitRotation(x float64, y float64, z float64) {
	identityMatrix := Matrix4f{M: [4][4]float64{}}
	identityMatrix.InitIdentity()

	rotX := identityMatrix
	rotY := identityMatrix
	rotZ := identityMatrix

	_x := x * (math.Pi / 180)
	_y := y * (math.Pi / 180)
	_z := z * (math.Pi / 180)

	rotZ.M[0][0] = math.Cos(_z)
	rotZ.M[0][1] = -math.Sin(_z)
	rotZ.M[1][0] = math.Sin(_z)
	rotZ.M[1][1] = math.Cos(_z)

	rotX.M[1][1] = math.Cos(_x)
	rotX.M[1][2] = -math.Sin(_x)
	rotX.M[2][1] = math.Sin(_x)
	rotX.M[2][2] = math.Cos(_x)

	rotY.M[0][0] = math.Cos(_y)
	rotY.M[0][2] = -math.Sin(_y)
	rotY.M[2][0] = math.Sin(_y)
	rotY.M[2][2] = math.Cos(_y)

	mulYX := rotY.Mul(&rotX)
	m.M = rotZ.Mul(&mulYX).M
}

func (m *Matrix4f) InitProjection(fov float64, width float64, height float64, zNear float64, zFar float64) {
	aspectRatio := width/height
	tanHalfFOV := math.Tan((fov/2) * (math.Pi/180))
	zRange := zNear - zFar

	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if i == j {
				if i == 0 {
					m.M[i][j] = 1 / (tanHalfFOV * aspectRatio)
				} else if i == 1 {
					m.M[i][j] = 1 / (tanHalfFOV)
				} else if i == 2 {
					m.M[i][j] = (-zNear - zFar)/zRange
				} else if i == 3 {
					m.M[i][j] = 0
				}
			} else {
				if i == 3 && j == 2 {
					m.M[i][j] = 1
				} else if i == 2 && j == 3 {
					m.M[i][j] = (2 * zFar * zNear) / zRange
				} else {
					m.M[i][j] = 0
				}
			}
		}
	}
}

func (m *Matrix4f) Mul(r *Matrix4f) Matrix4f {
	matrixRes := Matrix4f{}

	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			matrixRes.M[i][j] = m.M[i][0]*r.M[0][j] +
				m.M[i][1]*r.M[1][j] +
				m.M[i][2]*r.M[2][j] +
				m.M[i][3]*r.M[3][j]

		}
	}
	return matrixRes
}

func (m *Matrix4f) flatten() []float32 {
	matrixSlice := make([]float32, 0)
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			matrixSlice = append(matrixSlice, float32(m.M[i][j]))
		}
	}
	return matrixSlice
}
