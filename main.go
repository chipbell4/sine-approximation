package main

import (
	"fmt"
	"math"
)

func main() {
	// Print an error table for sine calculations
	for i := 0; i < 201; i += 1 {
		// interpolate up to π / 4
		frac := float64(i) / 400
		theta := frac * math.Pi

		// calculate error
		fast := FastSin(theta)
		orig := math.Sin(theta)
		percentage_error := (math.Abs(fast-orig) / orig) * 100

		fmt.Printf("%f π\t%.2f%%\n", frac, percentage_error)
	}
}
