package data

import (
	"math"
	"time"
)

type DataPoint struct {
	X float64
	Y float64
}

type Metric struct {
	Name   string
	Points []DataPoint
}

type Chart struct {
	XMin    float64
	XMax    float64
	YMin    float64
	YMax    float64
	Metrics []Metric
}

func generateSinChart(xMin float64, xMax float64, n int) *Chart {
	step := (xMax - xMin) / float64(n-1)

	points := make([]DataPoint, n)

	for i := 0; i < n; i++ {
		x := xMin + float64(i)*step
		points[i] = DataPoint{x, float64(math.Sin(float64(x)))}
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

var step = float64(0.0)

var points = make([]DataPoint, 0)

func getCurrentTimestamp() float64 {
	return float64(time.Now().UnixNano() / int64(time.Millisecond))
}

func AddDataPoint(value float64) {
	if len(points) == 0 {
		points = append(points, DataPoint{getCurrentTimestamp(), value})
	} else {
		points = append(points, DataPoint{getCurrentTimestamp(), 0.8*value + 0.2*points[len(points)-1].Y})
	}
}

func GetChart() *Chart {
	now := getCurrentTimestamp()
	start := now - 10001
	end := now - 1000

	trim := -1
	for trim+1 < len(points) && points[trim+1].X < start {
		trim++
	}

	if trim != -1 {
		points = points[trim:]
	}

	return &Chart{
		start,
		end,
		0,
		500,
		[]Metric{
			Metric{"tx_bytes", points},
		},
	}
}
