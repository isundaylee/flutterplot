package main

import (
	"net"
	"runtime"

	"github.com/isundaylee/flutterplot/graphics"
	"github.com/isundaylee/flutterplot/transport"
)

func main() {
	runtime.LockOSThread()

	// Init the graphics part
	if err := graphics.Init(); err != nil {
		panic(err)
	}

	// Connect to flutter server
	addr, err := net.ResolveTCPAddr("tcp", "localhost:4999")
	if err != nil {
		panic(err)
	}
	transport.Connect(addr)

	// Main loop
	for !graphics.ShouldExit() {
		graphics.Render()
	}

	// Cleanup
	graphics.Cleanup()
}
