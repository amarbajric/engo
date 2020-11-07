package core

import "math"

var (
	vertices = []Vertex{
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
		{Vector3f{
			X: 0,
			Y: -1,
			Z: 1,
		}},
	}
	indices = []uint32{0,1,3,3,1,2,2,1,0,0,2,3}
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
	}, scale: Vector3f{
		X: 1,
		Y: 1,
		Z: 1,
	}}
	tmp = 0.0
	diamondMesh Mesh
)

func start() {
	shader.Init()
	shader.AddVertexShader(loadShader("vertexShader.glsl"))
	shader.AddFragmentShader(loadShader("fragmentShader.glsl"))
	shader.CompileShader()
	//mesh.AddVertices(vertices, indices)

	shader.addUniform("transform")

	diamondMesh = loadMesh("res/models/pyramid.obj")
}

func update() {
	tmp += delta

	sinTmp := math.Sin(tmp)
	transform.setTranslation1f(math.Sin(tmp), 0, 0)
	transform.setRotation1f(0, math.Sin(tmp) * 180, 0)
	transform.setScale1f(sinTmp, sinTmp, sinTmp)
}

func render() {
	shader.Bind()
	shader.setUniform4m("transform", transform.getTransformation())
	//mesh.Draw()
	diamondMesh.Draw()
}
