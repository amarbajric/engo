package core

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Mesh struct {
	vbo  uint32
	size int32
}

func (m *Mesh) AddVertices(vertices []Vertex) {
	m.size = int32(len(vertices) * VERTEX_SIZE)
	gl.GenBuffers(1, &m.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4 * int(m.size), gl.Ptr(createVertexSlice(vertices)), gl.STATIC_DRAW)
}

func (m *Mesh) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.GenVertexArrays(1, &m.vbo)
	gl.BindVertexArray(m.vbo)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.DrawArrays(gl.TRIANGLES, 0, m.size)
	gl.DisableVertexAttribArray(0)

	window.SwapBuffers()
	glfw.PollEvents()
}

func createVertexSlice(vertices []Vertex) []float32 {
	vertexSlice := make([]float32, 0)
	for _, vertex := range vertices {
		vertexSlice = append(vertexSlice, float32(vertex.Pos.X), float32(vertex.Pos.Y), float32(vertex.Pos.Z))
	}
	return vertexSlice
}