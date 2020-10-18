package main

import (
	"engo/core"
	"runtime"
)

func main() {
	runtime.LockOSThread()
	core.Init()
	defer core.Stop()
	core.Start()
}
