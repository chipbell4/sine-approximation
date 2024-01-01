package main

import "math"

func OriginalSine(x float64) float64 {
	return math.Sin(x)
}

func FastSin(x float64) float64 {
	return x - x*x*x/6
}
