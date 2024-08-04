package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApproximatePI(t *testing.T) {
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

func BenchmarkApproximatePI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		approximatePI(10000)
	}
}
