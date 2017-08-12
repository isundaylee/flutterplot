package graphics

import (
	"image/color"

	"github.com/isundaylee/flutterplot/data"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	logging "github.com/op/go-logging"
)

const (
	windowWidth  = 960
	windowHeight = 540

	chartPadding = 40.0
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

func windowPoint(x float32, y float32) [3]float32 {
	return [3]float32{
		(x - windowWidth/2) / (windowWidth / 2),
		(y - windowHeight/2) / (windowHeight / 2),
		0.0,
	}
}

func drawAxes() {
	gl.LineWidth(10.0)

	radius := float32(0.5)

	// Draw the X axis
	drawPrimitive(gl.TRIANGLE_FAN, [][3]float32{
		windowPoint(chartPadding-radius, chartPadding-radius),
		windowPoint(chartPadding-radius, chartPadding+radius),
		windowPoint(windowWidth-chartPadding+radius, chartPadding+radius),
		windowPoint(windowWidth-chartPadding+radius, chartPadding-radius),
	})

	// Draw the Y axis
	drawPrimitive(gl.TRIANGLE_FAN, [][3]float32{
		windowPoint(chartPadding-radius, chartPadding-radius),
		windowPoint(chartPadding+radius, chartPadding-radius),
		windowPoint(chartPadding+radius, windowHeight-chartPadding+radius),
		windowPoint(chartPadding-radius, windowHeight-chartPadding+radius),
	})
}

func chartPoint(chart *data.Chart, x float32, y float32) [3]float32 {
	xRatio := float32((windowWidth - 2*chartPadding) / windowWidth)
	yRatio := float32((windowHeight - 2*chartPadding) / windowHeight)

	return [3]float32{
		xRatio * (x - (chart.XMax+chart.XMin)/2) / ((chart.XMax - chart.XMin) / 2),
		yRatio * (y - (chart.YMax+chart.YMin)/2) / ((chart.YMax - chart.YMin) / 2),
		0.0,
	}
}

func drawChart() {
	chart := data.GetChart()

	for _, metric := range chart.Metrics {
		points := make([][3]float32, 2*len(metric.Points))
		for i, point := range metric.Points {
			points[2*i] = chartPoint(chart, point.X, chart.YMin)
			points[2*i+1] = chartPoint(chart, point.X, point.Y)
		}

		drawPrimitive(gl.TRIANGLE_STRIP, points)
	}
}

func Render() {
	gl.ClearColor(
		float32(backgroundColor.R)/255.0,
		float32(backgroundColor.G)/255.0,
		float32(backgroundColor.B)/255.0,
		float32(backgroundColor.A)/255.0,
	)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	drawChart()
	drawAxes()

	glfw.PollEvents()
	window.SwapBuffers()
}

func ShouldExit() bool {
	return window.ShouldClose()
}

func Cleanup() {
	glfw.Terminate()
}
