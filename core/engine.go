package core

import (
	"fmt"
	"time"

	// Use 4.1 as macOS has deprecated OpenGL since 10.15 and only supports up to v4.1
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 800
	height = 600
	title = "3D Engine"
	frameCap = 5000.0
)

var (
	window *glfw.Window
	program uint32
	isRunning = false
	delta float64
)

func Init() {
	window = InitGlfw()
	program = InitOpenGL()
}

func Start() {
	if isRunning {
		return
	}
	run()
}

func Stop() {
	if !isRunning {
		return
	}
	isRunning = false
}

func run() {
	isRunning = true

	start()

	frames := 0
	frameCounter := int64(0)
	frameTime := 1.0 / frameCap
	unprocessedTime := 0.0
	lastTime := time.Now()

	for isRunning {
		shouldRender := false
		passedTime := time.Since(lastTime)
		startTime := time.Now()
		lastTime = startTime

		unprocessedTime += float64(passedTime.Nanoseconds()) / float64(time.Second.Nanoseconds())
		frameCounter += passedTime.Nanoseconds()

		for unprocessedTime > frameTime {
			shouldRender = true
			unprocessedTime -= frameTime

			if window.ShouldClose() {
				Stop()
			}

			delta = frameTime
			//TODO Update Game
			update()

			if frameCounter >= time.Second.Nanoseconds() {
				fmt.Println(frames, "FPS")
				frames = 0
				frameCounter = 0
			}
		}
		if shouldRender {
			render()
			frames++
		} else {
			time.Sleep(time.Millisecond)
		}
	}
	cleanUp()
}

func cleanUp() {
	glfw.Terminate()
}

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)
	window.SwapBuffers()
	glfw.PollEvents()
}