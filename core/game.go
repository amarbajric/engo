package core

var data = []Vertex{
	{Vector3f{
		X: .05,
		Y: .05,
		Z: 0,
	}},
	{Vector3f{
		X: -.05,
		Y: -.05,
		Z: 0,
	}},
	{Vector3f{
		X: .05,
		Y: -.05,
		Z: 0,
	}},
}
var mesh = Mesh{}

func start() {
	mesh.AddVertices(data)
}

func update() {
}

func render() {
	//draw()
	mesh.Draw()
}
