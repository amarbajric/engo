package core

import "math"

var (
	data = []Vertex{
		{Vector3f{
			X: -1,
			Y: -1,
			Z: 0,
		}},
		{Vector3f{
			X: 0,
			Y: 1,
			Z: 0,
		}},
		{Vector3f{
			X: 1,
			Y: -1,
			Z: 0,
		}},
	}
	mesh = Mesh{}
	shader = Shader{}
	tmp = 0.0
)

func start() {
	shader.Init()
	shader.AddVertexShader(loadShader("vertexShader.glsl"))
	shader.AddFragmentShader(loadShader("fragmentShader.glsl"))
	shader.CompileShader()
	mesh.AddVertices(data)

	shader.addUniform("uniformFloat")
}

func update() {
	tmp += delta
	shader.setUniform1f("uniformFloat", float32(math.Abs(math.Sin(tmp))))
}

func render() {
	shader.Bind()
	mesh.Draw()
}
