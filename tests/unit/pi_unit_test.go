package unit

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPIConstants(t *testing.T) {
	assert.Equal(t, math.Pi, piConstant, "piConstant should be equal to math.Pi")
}

func TestPIApproximation(t *testing.T) {
	tests := []struct {
		name        string
		numPoints   int
		expectedPI float64
	}{
		{"Small", 100, 3.14},
		{"Medium", 1000, 3.141},
		{"Large", 10000, 3.1415},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualPI := approximatePI(tt.numPoints)
			assert.InDelta(t, tt.expectedPI, actualPI, 0.01)
		})
	}
}

func BenchmarkPIApproximation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		approximatePI(10000)
	}
}

func approximatePI(numPoints int) float64 {
	var sum float64
	for i := 0; i < numPoints; i++ {
		x := float64(i) / float64(numPoints-1)
		y := 4 * (1 - x*x)
		sum += y
	}
	return sum / float64(numPoints)
}

const piConstant = 3.14159265359
