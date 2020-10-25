package core

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

func GetKey(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {}

func GetMouseButton(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {}

func GetMousePosition(window *glfw.Window, xPos float64, yPos float64) {}
