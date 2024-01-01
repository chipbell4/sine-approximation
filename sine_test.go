package main

import (
	"testing"
)

func BenchmarkOriginalSine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OriginalSine(0.5)
	}
}

func BenchmarkFastSine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastSin(0.5)
	}
}
