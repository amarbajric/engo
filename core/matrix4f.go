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