package main

import (
	"runtime"

	"github.com/isundaylee/flutterplot/graphics"
)

func main() {
	runtime.LockOSThread()

	if err := graphics.Init(); err != nil {
		panic(err)
	}

	for !graphics.ShouldExit() {
		graphics.Render()
	}

	graphics.Cleanup()
}
