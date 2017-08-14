package graphics

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/isundaylee/flutterplot/logger"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	vertexShaderSource = `
		#version 410

		in vec3 vp;
        layout(location = 1) in vec3 color;
        out vec3 fragmentColor;

		void main() {
			gl_Position = vec4(vp, 1.0);
            fragmentColor = color;
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 410

        in vec3 fragmentColor;
		out vec3 color;

		void main() {
			color = fragmentColor;
		}
	` + "\x00"

	primitiveBufferPointCount = 1000
)

func initGlfw() (*glfw.Window, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Samples, 8)

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
	logger.Logger.Info("OpenGL version is " + version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.Enable(gl.MULTISAMPLE)
	gl.LinkProgram(prog)

	initBuffers()

	return prog, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

type Primitive struct {
	mode  uint32
	begin int32
	end   int32
}

var pointBuffer [][3]float32
var colorBuffer [][3]float32
var pointsUsed = int32(0)

var vboPoint uint32
var vboColor uint32
var vao uint32

var primitives []Primitive

func initBuffers() {
	pointBuffer = make([][3]float32, primitiveBufferPointCount)
	colorBuffer = make([][3]float32, primitiveBufferPointCount)

	gl.GenBuffers(1, &vboPoint)
	gl.GenBuffers(1, &vboColor)

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vboPoint)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, vboColor)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)
}

func PreRender() {
	pointsUsed = 0
	primitives = make([]Primitive, 0)
}

func PostRender() {
	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vboPoint)
	gl.BufferData(gl.ARRAY_BUFFER, 4*3*len(pointBuffer), gl.Ptr(pointBuffer), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ARRAY_BUFFER, vboColor)
	gl.BufferData(gl.ARRAY_BUFFER, 4*3*len(colorBuffer), gl.Ptr(colorBuffer), gl.STATIC_DRAW)

	for _, segment := range primitives {
		gl.DrawArrays(segment.mode, segment.begin, segment.end-segment.begin)
	}
}

func drawPrimitive(mode uint32, points [][3]float32, color color.RGBA) {
	begin := pointsUsed

	for i := range points {
		pointBuffer[pointsUsed] = points[i]

		colorBuffer[pointsUsed] = [3]float32{
			float32(color.R) / 255.0,
			float32(color.G) / 255.0,
			float32(color.B) / 255.0,
		}

		pointsUsed++
		if pointsUsed >= primitiveBufferPointCount {
			panic("OpenGL point buffer full.")
		}
	}

	end := pointsUsed

	primitives = append(primitives, Primitive{mode, begin, end})
}
