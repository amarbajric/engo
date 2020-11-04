package core

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Mesh struct {
	vbo  uint32
	ibo  uint32
	vboSize int32
	iboSize int32
}

func (m *Mesh) AddVertices(vertices []Vertex, indices []uint32) {
	m.vboSize = int32(len(vertices) * VERTEX_SIZE)
	m.iboSize = int32(len(indices))

	// create a Vector Buffer Object that will store the vertices on video memory
	gl.GenBuffers(1, &m.vbo)
	// allocate space and upload the data from CPU to GPU
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*int(m.vboSize), gl.Ptr(createVertexSlice(vertices)), gl.STATIC_DRAW)

	// create an element array buffer in video memory
	// that will store the indices array and set the pointer under 'ibo'
	gl.GenBuffers(1, &m.ibo)
	// transfer the data from indices slice to ibo in video memory
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*int(m.iboSize), gl.Ptr(indices), gl.STATIC_DRAW)
}

func (m *Mesh) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.GenVertexArrays(1, &m.vbo)
	gl.BindVertexArray(m.vbo)
	gl.EnableVertexAttribArray(0)
	// specify data location for openGL
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ibo)
	// draw the elements (i.e. vertices that are already stored in video memory) in the given order of ibo
	gl.DrawElements(gl.TRIANGLES, m.iboSize, gl.UNSIGNED_INT, gl.PtrOffset(0))

	// inactivates the generic vertex attribute data of index 0
	gl.DisableVertexAttribArray(0)

	// swap active frame with a newly rendered frame
	window.SwapBuffers()
	// get all events from the user (e.g. input)
	glfw.PollEvents()
}

func createVertexSlice(vertices []Vertex) []float32 {
	vertexSlice := make([]float32, 0)
	for _, vertex := range vertices {
		vertexSlice = append(vertexSlice, float32(vertex.Pos.X), float32(vertex.Pos.Y), float32(vertex.Pos.Z))
	}
	return vertexSlice
}
