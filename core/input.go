package core

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)


func GetKey(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	movAmt := 10 * delta
	rotAmt := 100 * delta

	if key == glfw.KeyW {
		transform.camera.Move(transform.camera.forward, movAmt)
	} else if key == glfw.KeyS {
		transform.camera.Move(transform.camera.forward, -movAmt)
	} else if key == glfw.KeyA {
		transform.camera.Move(transform.camera.GetLeft(), movAmt)
	} else if key == glfw.KeyD {
		transform.camera.Move(transform.camera.GetRight(), movAmt)
	} else if key == glfw.KeyUp {
		transform.camera.RotateX(-rotAmt)
	} else if key == glfw.KeyDown {
		transform.camera.RotateX(rotAmt)
	} else if key == glfw.KeyLeft {
		transform.camera.RotateY(-rotAmt)
	} else if key == glfw.KeyRight {
		transform.camera.RotateY(rotAmt)
	}
}

func GetMouseButton(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
}

func GetMousePosition(window *glfw.Window, xPos float64, yPos float64) {}
