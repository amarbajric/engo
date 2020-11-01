package core

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
)

func start() {
	shader.Init()
	shader.AddVertexShader(loadShader("vertexShader.glsl"))
	shader.AddFragmentShader(loadShader("fragmentShader.glsl"))
	shader.CompileShader()
	mesh.AddVertices(data)
}

func update() {
}

func render() {
	shader.Bind()
	mesh.Draw()
}
