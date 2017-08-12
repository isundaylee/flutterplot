package data

import (
	"math"
)

type DataPoint struct {
	X float32
	Y float32
}

type Metric struct {
	Name   string
	Points []DataPoint
}

type Chart struct {
	XMin    float32
	XMax    float32
	YMin    float32
	YMax    float32
	Metrics []Metric
}

func generateSinChart(xMin float32, xMax float32, n int) *Chart {
	step := (xMax - xMin) / float32(n-1)

	points := make([]DataPoint, n)

	for i := 0; i < n; i++ {
		x := xMin + float32(i)*step
		points[i] = DataPoint{x, float32(math.Sin(float64(x)))}
	}

	return &Chart{
		xMin,
		xMax,
		-1.0,
		1.0,
		[]Metric{
			Metric{"sin", points},
		},
	}
}

var step = float32(0.0)

func GetChart() *Chart {
	step += 0.01
	return generateSinChart(step, step+3.0, 100)
}
