package graphics

import (
	"image/color"

	"github.com/isundaylee/flutterplot/data"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	windowWidth  = 960
	windowHeight = 540

	chartPadding = 40.0
)

var backgroundColor = color.RGBA{20, 20, 20, 255}
var axesColor = color.RGBA{52, 152, 219, 255}
var chartColor = color.RGBA{52, 152, 219, 255}

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

	radius := float32(1)

	// Draw the X axis
	drawPrimitive(gl.TRIANGLE_FAN, [][3]float32{
		windowPoint(chartPadding-radius, chartPadding-radius),
		windowPoint(chartPadding-radius, chartPadding+radius),
		windowPoint(windowWidth-chartPadding+radius, chartPadding+radius),
		windowPoint(windowWidth-chartPadding+radius, chartPadding-radius),
	}, axesColor)

	// Draw the Y axis
	drawPrimitive(gl.TRIANGLE_FAN, [][3]float32{
		windowPoint(chartPadding-radius, chartPadding-radius),
		windowPoint(chartPadding+radius, chartPadding-radius),
		windowPoint(chartPadding+radius, windowHeight-chartPadding+radius),
		windowPoint(chartPadding-radius, windowHeight-chartPadding+radius),
	}, axesColor)
}

func bound(x float64, lower float64, upper float64) float64 {
	if x < lower {
		return lower
	}

	if x > upper {
		return upper
	}

	return x
}

func chartPoint(chart *data.Chart, x float64, y float64) [3]float32 {
	xRatio := float64((windowWidth - 2*chartPadding) / windowWidth)
	yRatio := float64((windowHeight - 2*chartPadding) / windowHeight)

	return [3]float32{
		float32(xRatio * bound((x-(chart.XMax+chart.XMin)/2)/((chart.XMax-chart.XMin)/2), -1, 1)),
		float32(yRatio * bound((y-(chart.YMax+chart.YMin)/2)/((chart.YMax-chart.YMin)/2), -1, 1)),
		0.0,
	}
}

func drawChart() {
	chart := data.GetChart()

	for _, metric := range chart.Metrics {
		if len(metric.Points) == 0 {
			continue
		}

		points := make([][3]float32, 2*(len(metric.Points)+1))
		// points[0] = chartPoint(chart, chart.XMin, chart.YMin)
		// points[1] = chartPoint(chart, chart.XMin, metric.Points[0].Y)
		for i, point := range metric.Points {
			points[2*i+0] = chartPoint(chart, point.X, chart.YMin)
			points[2*i+1] = chartPoint(chart, point.X, point.Y)
		}
		points[2*len(metric.Points)+0] = chartPoint(chart, chart.XMax, chart.YMin)
		points[2*len(metric.Points)+1] = chartPoint(chart, chart.XMax, metric.Points[len(metric.Points)-1].Y)

		drawPrimitive(gl.TRIANGLE_STRIP, points, chartColor)
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
