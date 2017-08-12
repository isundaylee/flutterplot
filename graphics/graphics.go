package graphics

import (
	"image/color"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	logging "github.com/op/go-logging"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

var backgroundColor = color.RGBA{20, 20, 20, 255}

var log = logging.MustGetLogger("example")

var window *glfw.Window
var program uint32

func Init() error {
	var err error

	window, err = initGlfw()
	if err != nil {
		return err
	}

	program, err = initOpenGL()
	if err != nil {
		return err
	}

	return nil
}

func initGlfw() (*glfw.Window, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	var err error
	window, err = glfw.CreateWindow(windowWidth, windowHeight, "Flutter Plot", nil, nil)
	if err != nil {
		return nil, err
	}

	window.MakeContextCurrent()

	return window, nil
}

func initOpenGL() (uint32, error) {
	if err := gl.Init(); err != nil {
		return 0, err
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Info("OpenGL version is " + version)

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog, nil
}

func draw(window *glfw.Window, program uint32) {
	gl.ClearColor(
		float32(backgroundColor.R)/255.0,
		float32(backgroundColor.G)/255.0,
		float32(backgroundColor.B)/255.0,
		float32(backgroundColor.A)/255.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	glfw.PollEvents()
	window.SwapBuffers()
}

func Render() {
	draw(window, program)
}

func ShouldExit() bool {
	return window.ShouldClose()
}

func Cleanup() {
	glfw.Terminate()
}
