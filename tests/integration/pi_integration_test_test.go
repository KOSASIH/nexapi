package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePI(t *testing.T) {
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
			pi, err := CalculatePI(context.Background(), tt.numPoints)
			assert.NoError(t, err)

			assert.InDelta(t, tt.expectedPI, pi, 0.01)
		})
	}
}

func BenchmarkCalculatePI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CalculatePI(context.Background(), 10000)
		if err!= nil {
			b.Errorf("Expected CalculatePI to succeed, but got error: %s", err)
		}
	}
}
