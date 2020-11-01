package core

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

func (m *Matrix4f) Mul(r *Matrix4f) Matrix4f {
	matrixRes := Matrix4f{}

	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			matrixRes.M[i][j] = m.M[i][0] * r.M[0][j] +
								m.M[i][1] * r.M[1][j] +
								m.M[i][2] * r.M[2][j] +
								m.M[i][3] * r.M[3][j]


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