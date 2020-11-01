package core

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"log"
)

// initializes OpenGL and returns an initialized program
func InitOpenGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.FrontFace(gl.CW)
	gl.CullFace(gl.BACK)
	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.DEPTH_TEST)

	//TODO: Depth clamp

	gl.Enable(gl.FRAMEBUFFER_SRGB)
}
