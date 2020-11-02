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
	transform = Transform{translation: Vector3f{
		X: 0,
		Y: 0,
		Z: 0,
	}, rotation: Vector3f{
		X: 0,
		Y: 0,
		Z: 0,
	}}
	tmp = 0.0
)

func start() {
	shader.Init()
	shader.AddVertexShader(loadShader("vertexShader.glsl"))
	shader.AddFragmentShader(loadShader("fragmentShader.glsl"))
	shader.CompileShader()
	mesh.AddVertices(data)

	shader.addUniform("transform")
}

func update() {
	tmp += delta
	transform.setTranslation1f(math.Sin(tmp), 0, 0)
	transform.setRotation1f(0, 0, math.Sin(tmp) * 180)
}

func render() {
	shader.Bind()
	shader.setUniform4m("transform", transform.getTransformation())
	mesh.Draw()
}
