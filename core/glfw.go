package core

import "github.com/go-gl/glfw/v3.3/glfw"

// initGlfw initializes glfw and returns a Window to use.
func InitGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}
	// bind window to current thread
	window.MakeContextCurrent()

	return window
}

